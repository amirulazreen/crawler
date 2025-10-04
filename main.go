package main

import (
	"fmt"

	// chip "github.com/amirulazreen/chip-crawler/src"
	libraries "github.com/amirulazreen/chip-crawler/libraries"
	models "github.com/amirulazreen/chip-crawler/libraries/models"
)

func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <website_url>")
	// 	os.Exit(1)
	// }

	// website := os.Args[1]

	// Option 1: Run the crawler (commented out)
	// chip.Crawler(website)

	// Option 2: Generate text using OpenAI
	//

	texts := `Our ProductsCOLLECTPayment ServicesCONTROLExpense ManagementSendPayouts in real-timeExpenseTeam's Expenses ManagementADVANCEPay-As-You-Sell Advance™COMPLIANCERisk ManagementCOINTreasury ManagementBlogAPICareersContact usContact usLog inStart Now CHIP AdvanceIN PARTNERSHIP WITHGrow Your Sales with CHIP AdvanceAn exclusive Shariah-based financing solution for your business, offering a frictionless line of credit that's accessible at all times.Contact our teamHow CHIP Advance works?STEP 1Fill in the form to unlock the offer.STEP 2Submit an advance funds request.STEP 3Receive funds within 48 hours after approval.STEP 4Use the funds to grow your business and repay as your revenue grows.Pay-as-you-sell advance™Access Funds Up to 10x Your Weekly SalesGet instant access to funds based on your sales, with flexible repayment linked directly to your business performance.Repay As You EarnAutomatically Repay from Your Sales, Keeping Cash Flow SimpleEnjoy automatic, hassle-free repayments with no interest or late fees, making cash flow management effortless.Empower Your GrowthFocus on Growing Your Business, Not Managing FinancesWith no interest or late fees, CHIP Advance lets you concentrate on scaling your business, while we take care of the financing.ACHIEVEMENTS UNLOCKED!Total Sales Advance Fund Disbursed To MerchantsMany merchants have leveraged CHIP Advance to access flexible funding and grow their businesses.RM5.0 millionMMAs of August 2025Why use CHIP Advance?No Interest, No Late FeesGet an advance with a one-time fee of 6% and no additional interest charges or late fees, ever.Funds in 48 HoursSubmit your advance request in under 5 minutes and receive funds within 48 hours.Automated, Flexible RepaymentsAutomatically repay through future sales via your settlement partner.Advance Funds Based on SalesYour available advance increases as your sales grow.Certified Shariah-CompliantSeedflex’s Pay-As-You-Sell Advance™️ is certified as Shariah-compliant.Robust SecuritySSL-secured connection for your peace of mind.Frequently Asked QuestionsWho is eligible for CHIP Advance?Merchants must meet basic eligibility criteria, including a consistent sales history. Specific requirements may vary, so contact us to confirm your eligibility.Do I need collateral to qualify?No, CHIP Advance is unsecured financing, meaning you don’t need to provide any assets, like property or equipment, as collateral. Approval is based on your sales performance and business profile.How much funding can I get?The funding amount is determined by your historical sales data. You may qualify for an advance of up to 10 weeks’ worth of sales.Are there any hidden fees or charges?There are no hidden fees. CHIP Advance charges a one-time fee of 6%, with no interest or late fees.How does repayment work?Repayment is fully automated and tied to your sales. A fixed percentage of your daily transactions is deducted until the advance is fully repaid.What happens if my sales slow down?CHIP Advance is designed to adjust with your revenue. If your sales slow down, repayments decrease, minimizing the impact on your business.Register Your InterestFill in the interest form to allow CHIP to share your Company Information with Seedflex to access your available amount.You will receive your login details from Seedflex within 2 business days.Full name*Organization / Company name*Email*Contact number*By checking the box, I consent for CHIP to share my personal information with Seedflex*SubmitCHIP IN SDN. BHD. (202201010914 (1456611-H))Lot 3A-01A, Level 3AGlo Damansara Shopping Mall,699, Jalan Damansara, Taman Tun Dr Ismail60000 Kuala Lumpur, MalaysiaOffice line: (+60)3 2935 9253CHIP ProductsCOLLECTPayment ServicesSENDPayouts in real-timeEXPENSETeam Expense ManagementADVANCEPay-As-You-Sell Advance™COMPLIANCERisk ManagementCOINTreasury ManagementPartnersBecome a referral partnerReferrer LoginResourcesBlogAPICHIP Services StatusFPX Bank StatusCompanyInvestor RelationsContact usCareersBrand assetsDownload wallpaperFollow usTerms of Service · Privacy Policy · © 2025 CHIPChat with us{"props":{"pageProps":{"_sentryTraceData":"b4a867f20bbda72ac20886fa7d1f7369-678177c02a3afe0b-1","_sentryBaggage":"sentry-environment=production,sentry-release=jDsTDFf4g7e7OYWfH40gQ,sentry-public_key=55f8e3ef4da14d0501741236e07d0229,sentry-trace_id=b4a867f20bbda72ac20886fa7d1f7369,sentry-sample_rate=1,sentry-transaction=GET%20%2Fcontrol%2Fadvance,sentry-sampled=true"}},"page":"/control/advance","query":{},"buildId":"jDsTDFf4g7e7OYWfH40gQ","isFallback":false,"isExperimentalCompile":false,"appGip":true,"scriptLoader":[]}`

	prompt := []models.Message{
		{
			Role: "system",
			Content: `You are a website content risk analyzer. For each page,
					you will summarize it and rank risk level: high (gambling, alcohol, adult), medium (marketing, crypto), low (safe, educational).
					Return JSON with fields: summary, topics, and risk_level.`,
		},
		{
			Role:    "user",
			Content: fmt.Sprintf("Analyze this webpage:\n\n%s", texts),
		},
	}

	param := models.Request{
		Model:       openAIOSS,
		Temperature: 0.2,
		Messages:    prompt,
	}

	text, err := libraries.GenerateText(param)
	if err != nil {
		fmt.Printf("Error generating text: %v\n", err)
		return
	}
	fmt.Printf("Generated text: %s\n", text)
}
