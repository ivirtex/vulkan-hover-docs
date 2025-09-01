# VK\_KHR\_load\_store\_op\_none(3) Manual Page

## Name

VK\_KHR\_load\_store\_op\_none - device extension



## [](#_registered_extension_number)Registered Extension Number

527

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Shahbaz Youssefi [\[GitHub\]syoussefi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_load_store_op_none%5D%20%40syoussefi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_load_store_op_none%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_load\_store\_op\_none](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_load_store_op_none.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-05-16

**Contributors**

- Shahbaz Youssefi, Google
- Bill Licea-Kane, Qualcomm Technologies, Inc.
- Tobias Hector, AMD

## [](#_description)Description

This extension provides `VK_ATTACHMENT_LOAD_OP_NONE_KHR` and `VK_ATTACHMENT_STORE_OP_NONE_KHR`, which are identically promoted from the `VK_EXT_load_store_op_none` extension.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_LOAD_STORE_OP_NONE_EXTENSION_NAME`
- `VK_KHR_LOAD_STORE_OP_NONE_SPEC_VERSION`
- Extending [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html):
  
  - `VK_ATTACHMENT_LOAD_OP_NONE_KHR`
- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html):
  
  - `VK_ATTACHMENT_STORE_OP_NONE_KHR`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the KHR suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

Note

While `VK_ATTACHMENT_STORE_OP_NONE` is part of Vulkan 1.3, this extension was not promoted to core Vulkan 1.3 either in whole or in part. This functionality was promoted from `VK_KHR_dynamic_rendering`.

## [](#_version_history)Version History

- Revision 1, 2023-05-16 (Shahbaz Youssefi)
  
  - Initial revision, based on VK\_EXT\_load\_store\_op\_none.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_load_store_op_none).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0