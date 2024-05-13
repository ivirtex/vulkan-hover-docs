# VK_KHR_load_store_op_none(3) Manual Page

## Name

VK_KHR_load_store_op_none - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

527

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_load_store_op_none%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_load_store_op_none%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_load_store_op_none](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_load_store_op_none.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-05-16

**Contributors**  
- Shahbaz Youssefi, Google

- Bill Licea-Kane, Qualcomm Technologies, Inc.

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension provides `VK_ATTACHMENT_LOAD_OP_NONE_KHR` and
`VK_ATTACHMENT_STORE_OP_NONE_KHR`, which are identically promoted from
the [`VK_EXT_load_store_op_none`](VK_EXT_load_store_op_none.html)
extension.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_LOAD_STORE_OP_NONE_EXTENSION_NAME`

- `VK_KHR_LOAD_STORE_OP_NONE_SPEC_VERSION`

- Extending [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html):

  - `VK_ATTACHMENT_LOAD_OP_NONE_KHR`

- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html):

  - `VK_ATTACHMENT_STORE_OP_NONE_KHR`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>While <code>VK_ATTACHMENT_STORE_OP_NONE</code> is part of Vulkan 1.3,
this extension was not promoted to core either in whole or in part. This
functionality was promoted from <a
href="VK_KHR_dynamic_rendering.html"><code>VK_KHR_dynamic_rendering</code></a>.</p></td>
</tr>
</tbody>
</table>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-05-16 (Shahbaz Youssefi)

  - Initial revision, based on VK_EXT_load_store_op_none.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_load_store_op_none"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
