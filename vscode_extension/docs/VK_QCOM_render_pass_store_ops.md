# VK\_QCOM\_render\_pass\_store\_ops(3) Manual Page

## Name

VK\_QCOM\_render\_pass\_store\_ops - device extension



## [](#_registered_extension_number)Registered Extension Number

302

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_render_pass_store_ops%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_render_pass_store_ops%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-03-25

**Contributors**

- Bill Licea-Kane, Qualcomm Technologies, Inc.

## [](#_description)Description

Render pass attachments **can** be read-only for the duration of a render pass.

Examples include input attachments and depth attachments where depth tests are enabled but depth writes are not enabled.

In such cases, there **can** be no contents generated for an attachment within the render area.

This extension adds a new [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) `VK_ATTACHMENT_STORE_OP_NONE_QCOM` specifying that the contents within the render area **may** not be written to memory, but that the prior contents of the attachment in memory are preserved. However, if any contents were generated within the render area during rendering, the contents of the attachment will be undefined inside the render area.

Note

The [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) `VK_ATTACHMENT_STORE_OP_STORE` **may** force an implementation to assume that the attachment was written and force an implementation to flush data to memory or to a higher level cache. The [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) `VK_ATTACHMENT_STORE_OP_NONE_QCOM` **may** allow an implementation to assume that the attachment was not written and allow an implementation to avoid such a flush.

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_RENDER_PASS_STORE_OPS_EXTENSION_NAME`
- `VK_QCOM_RENDER_PASS_STORE_OPS_SPEC_VERSION`
- Extending [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html):
  
  - `VK_ATTACHMENT_STORE_OP_NONE_QCOM`

## [](#_version_history)Version History

- Revision 1, 2019-12-20 (wwlk)
  
  - Initial version
- Revision 2, 2020-03-25 (wwlk)
  
  - Minor renaming

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_render_pass_store_ops)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0