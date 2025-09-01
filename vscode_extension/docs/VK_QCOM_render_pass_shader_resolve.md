# VK\_QCOM\_render\_pass\_shader\_resolve(3) Manual Page

## Name

VK\_QCOM\_render\_pass\_shader\_resolve - device extension



## [](#_registered_extension_number)Registered Extension Number

172

## [](#_revision)Revision

4

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_render_pass_shader_resolve%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_render_pass_shader_resolve%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-11-07

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

None.

**Contributors**

- Srihari Babu Alla, Qualcomm
- Bill Licea-Kane, Qualcomm
- Jeff Leger, Qualcomm

## [](#_description)Description

This extension allows a shader resolve to replace fixed-function resolve.

Fixed-function resolve is limited in function to simple filters of multisample buffers to a single sample buffer.

Fixed-function resolve is more performance efficient and/or power efficient than shader resolve for such simple filters.

Shader resolve allows a shader writer to create complex, non-linear filtering of a multisample buffer in the last subpass of a subpass dependency chain.

This extension also provides a bit which **can** be used to enlarge a sample region dependency to a fragment region dependency, so that a framebuffer-region dependency **can** replace a framebuffer-global dependency in some cases.

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_RENDER_PASS_SHADER_RESOLVE_EXTENSION_NAME`
- `VK_QCOM_RENDER_PASS_SHADER_RESOLVE_SPEC_VERSION`
- Extending [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html):
  
  - `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`
  - `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`

## [](#_issues)Issues

1\) Should this extension be named render\_pass\_shader\_resolve?

**RESOLVED** Yes.

This is part of suite of small extensions to render pass.

Following the style guide, instead of following VK\_KHR\_create\_renderpass2.

2\) Should the VK\_SAMPLE\_COUNT\_1\_BIT be required for each pColorAttachment and the DepthStencilAttachent?

**RESOLVED** No.

While this may not be a common use case, and while most fixed-function resolve hardware has this limitation, there is little reason to require a shader resolve to resolve to a single sample buffer.

3\) Should a shader resolve subpass be the last subpass in a render pass?

**RESOLVED** Yes.

To be more specific, it should be the last subpass in a subpass dependency chain.

4\) Do we need the `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM` bit?

**RESOLVED** Yes.

This applies when an input attachment’s sample count is equal to `rasterizationSamples`. Further, if `sampleShading` is enabled (explicitly or implicitly) then `minSampleShading` **must** equal 0.0.

However, this bit may be set on any subpass, it is not restricted to a shader resolve subpass.

## [](#_version_history)Version History

- Revision 1, 2019-06-28 (wwlk)
  
  - Initial draft
- Revision 2, 2019-11-06 (wwlk)
  
  - General clean-up/spec updates
  - Added issues
- Revision 3, 2019-11-07 (wwlk)
  
  - Typos
  - Additional issues
  - Clarified that a shader resolve subpass is the last subpass in a subpass dependency chain
- Revision 4, 2020-01-06 (wwlk)
  
  - Change resolution of Issue 1 (*render\_pass*, not *renderpass*)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_render_pass_shader_resolve).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0