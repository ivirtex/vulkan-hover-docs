# VK_EXT_load_store_op_none(3) Manual Page

## Name

VK_EXT_load_store_op_none - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

401

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_load_store_op_none](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_load_store_op_none.html) extension

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_load_store_op_none%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_load_store_op_none%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-06-06

**Contributors**  
- Shahbaz Youssefi, Google

- Bill Licea-Kane, Qualcomm Technologies, Inc.

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension incorporates `VK_ATTACHMENT_STORE_OP_NONE_EXT` from
[`VK_QCOM_render_pass_store_ops`](VK_QCOM_render_pass_store_ops.html),
enabling applications to avoid unnecessary synchronization when an
attachment is not written during a render pass.

Additionally, `VK_ATTACHMENT_LOAD_OP_NONE_EXT` is introduced to avoid
unnecessary synchronization when an attachment is not used during a
render pass at all. In combination with
`VK_ATTACHMENT_STORE_OP_NONE_EXT`, this is useful as an alternative to
preserve attachments in applications that cannot decide if an attachment
will be used in a render pass until after the necessary pipelines have
been created.

## <a href="#_promotion_to_vk_khr_load_store_op_none" class="anchor"></a>Promotion to `VK_KHR_load_store_op_none`

All functionality in this extension is included in
[`VK_KHR_load_store_op_none`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_load_store_op_none.html), with the
suffix changed to KHR. The original enum names are still available as
aliases of the KHR functionality.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_LOAD_STORE_OP_NONE_EXTENSION_NAME`

- `VK_EXT_LOAD_STORE_OP_NONE_SPEC_VERSION`

- Extending [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html):

  - `VK_ATTACHMENT_LOAD_OP_NONE_EXT`

- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html):

  - `VK_ATTACHMENT_STORE_OP_NONE_EXT`

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

- Revision 1, 2021-06-06 (Shahbaz Youssefi)

  - Initial revision, based on VK_QCOM_render_pass_store_ops.

  - Added VK_ATTACHMENT_LOAD_OP_NONE_EXT.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_load_store_op_none"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
