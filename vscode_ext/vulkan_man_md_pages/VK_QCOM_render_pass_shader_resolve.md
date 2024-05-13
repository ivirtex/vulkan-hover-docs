# VK_QCOM_render_pass_shader_resolve(3) Manual Page

## Name

VK_QCOM_render_pass_shader_resolve - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

172

## <a href="#_revision" class="anchor"></a>Revision

4

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_render_pass_shader_resolve%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_render_pass_shader_resolve%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows a shader resolve to replace fixed-function
resolve.

Fixed-function resolve is limited in function to simple filters of
multisample buffers to a single sample buffer.

Fixed-function resolve is more performance efficient and/or power
efficient than shader resolve for such simple filters.

Shader resolve allows a shader writer to create complex, non-linear
filtering of a multisample buffer in the last subpass of a subpass
dependency chain.

This extension also provides a bit which **can** be used to enlarge a
sample region dependency to a fragment region dependency, so that a
framebuffer-region dependency **can** replace a framebuffer-global
dependency in some cases.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_RENDER_PASS_SHADER_RESOLVE_EXTENSION_NAME`

- `VK_QCOM_RENDER_PASS_SHADER_RESOLVE_SPEC_VERSION`

- Extending
  [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html):

  - `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`

  - `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension be named render_pass_shader_resolve?

**RESOLVED** Yes.

This is part of suite of small extensions to render pass.

Following the style guide, instead of following
VK_KHR_create_renderpass2.

2\) Should the VK_SAMPLE_COUNT_1_BIT be required for each
pColorAttachment and the DepthStencilAttachent?

**RESOLVED** No.

While this may not be a common use case, and while most fixed-function
resolve hardware has this limitation, there is little reason to require
a shader resolve to resolve to a single sample buffer.

3\) Should a shader resolve subpass be the last subpass in a render
pass?

**RESOLVED** Yes.

To be more specific, it should be the last subpass in a subpass
dependency chain.

4\) Do we need the `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`
bit?

**RESOLVED** Yes.

This applies when an input attachment’s sample count is equal to
`rasterizationSamples`. Further, if `sampleShading` is enabled
(explicitly or implicitly) then `minSampleShading` **must** equal 0.0.

However, this bit may be set on any subpass, it is not restricted to a
shader resolve subpass.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-06-28 (wwlk)

  - Initial draft

- Revision 2, 2019-11-06 (wwlk)

  - General clean-up/spec updates

  - Added issues

- Revision 3, 2019-11-07 (wwlk)

  - Typos

  - Additional issues

  - Clarified that a shader resolve subpass is the last subpass in a
    subpass dependency chain

- Revision 4, 2020-01-06 (wwlk)

  - Change resolution of Issue 1 (*render_pass*, not *renderpass*)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_render_pass_shader_resolve"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
