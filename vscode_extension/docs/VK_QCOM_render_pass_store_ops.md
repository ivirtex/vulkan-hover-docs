# VK_QCOM_render_pass_store_ops(3) Manual Page

## Name

VK_QCOM_render_pass_store_ops - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

302

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_render_pass_store_ops%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_render_pass_store_ops%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-03-25

**Contributors**  
- Bill Licea-Kane, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

Render pass attachments **can** be read-only for the duration of a
render pass.

Examples include input attachments and depth attachments where depth
tests are enabled but depth writes are not enabled.

In such cases, there **can** be no contents generated for an attachment
within the render area.

This extension adds a new
[VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html)
`VK_ATTACHMENT_STORE_OP_NONE_QCOM` specifying that the contents within
the render area **may** not be written to memory, but that the prior
contents of the attachment in memory are preserved. However, if any
contents were generated within the render area during rendering, the
contents of the attachment will be undefined inside the render area.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html">VkAttachmentStoreOp</a>
<code>VK_ATTACHMENT_STORE_OP_STORE</code> <strong>may</strong> force an
implementation to assume that the attachment was written and force an
implementation to flush data to memory or to a higher level cache. The
<a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html">VkAttachmentStoreOp</a>
<code>VK_ATTACHMENT_STORE_OP_NONE_QCOM</code> <strong>may</strong> allow
an implementation to assume that the attachment was not written and
allow an implementation to avoid such a flush.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_RENDER_PASS_STORE_OPS_EXTENSION_NAME`

- `VK_QCOM_RENDER_PASS_STORE_OPS_SPEC_VERSION`

- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html):

  - `VK_ATTACHMENT_STORE_OP_NONE_QCOM`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-12-20 (wwlk)

  - Initial version

- Revision 2, 2020-03-25 (wwlk)

  - Minor renaming

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_render_pass_store_ops"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
