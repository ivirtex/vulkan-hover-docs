# VK\_EXT\_load\_store\_op\_none(3) Manual Page

## Name

VK\_EXT\_load\_store\_op\_none - device extension



## [](#_registered_extension_number)Registered Extension Number

401

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_load\_store\_op\_none](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_load_store_op_none.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_load_store_op_none%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_load_store_op_none%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-06-06

**Contributors**

- Shahbaz Youssefi, Google
- Bill Licea-Kane, Qualcomm Technologies, Inc.
- Tobias Hector, AMD

## [](#_description)Description

This extension incorporates `VK_ATTACHMENT_STORE_OP_NONE_EXT` from `VK_QCOM_render_pass_store_ops`, enabling applications to avoid unnecessary synchronization when an attachment is not written during a render pass.

Additionally, `VK_ATTACHMENT_LOAD_OP_NONE_EXT` is introduced to avoid unnecessary synchronization when an attachment is not used during a render pass at all. In combination with `VK_ATTACHMENT_STORE_OP_NONE_EXT`, this is useful as an alternative to preserve attachments in applications that cannot decide if an attachment will be used in a render pass until after the necessary pipelines have been created.

## [](#_promotion_to_vk_khr_load_store_op_none)Promotion to `VK_KHR_load_store_op_none`

All functionality in this extension is included in `VK_KHR_load_store_op_none`, with the suffix changed to KHR. The original enum names are still available as aliases of the KHR functionality.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_LOAD_STORE_OP_NONE_EXTENSION_NAME`
- `VK_EXT_LOAD_STORE_OP_NONE_SPEC_VERSION`
- Extending [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html):
  
  - `VK_ATTACHMENT_LOAD_OP_NONE_EXT`
- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html):
  
  - `VK_ATTACHMENT_STORE_OP_NONE_EXT`

Note

While `VK_ATTACHMENT_STORE_OP_NONE` is part of Vulkan 1.3, this extension was not promoted to core either in whole or in part. This functionality was promoted from `VK_KHR_dynamic_rendering`.

## [](#_version_history)Version History

- Revision 1, 2021-06-06 (Shahbaz Youssefi)
  
  - Initial revision, based on VK\_QCOM\_render\_pass\_store\_ops.
  - Added VK\_ATTACHMENT\_LOAD\_OP\_NONE\_EXT.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_load_store_op_none)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0