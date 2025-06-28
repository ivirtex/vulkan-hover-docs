# VK\_QCOM\_rotated\_copy\_commands(3) Manual Page

## Name

VK\_QCOM\_rotated\_copy\_commands - device extension



## [](#_registered_extension_number)Registered Extension Number

334

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_rotated_copy_commands%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_rotated_copy_commands%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-12-13

**Interactions and External Dependencies**

- This extension interacts with `VK_KHR_swapchain`
- This extension interacts with `VK_KHR_surface`

**Contributors**

- Jeff Leger, Qualcomm Technologies, Inc.
- Matthew Netsch, Qualcomm Technologies, Inc.

## [](#_description)Description

This extension extends adds an optional rotation transform to copy commands [vkCmdBlitImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2KHR.html), [vkCmdCopyImageToBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2KHR.html) and [vkCmdCopyBufferToImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2KHR.html). When copying between two resources, where one resource contains rotated content and the other does not, a rotated copy may be desired. This extension may be used in combination with VK\_QCOM\_render\_pass\_transform which adds rotated render passes.

This extension adds an extension structure to the following commands: vkCmdBlitImage2KHR, vkCmdCopyImageToBuffer2KHR and vkCmdCopyBufferToImage2KHR

## [](#_issues)Issues

1\) What is an appropriate name for the added extension structure? The style guide says “Structures which extend other structures through the `pNext` chain should reflect the name of the base structure they extend.”, but in this case a single extension structure is used to extend three base structures (vkCmdBlitImage2KHR, vkCmdCopyImageToBuffer2KHR and vkCmdCopyBufferToImage2KHR). Creating three identical structures with unique names seemed undesirable.

**RESOLVED**: Deviate from the style guide for extension structure naming.

2\) Should this extension add a rotation capability to vkCmdCopyImage2KHR?

**RESOLVED**: No. Use of rotated vkCmdBlitImage2KHR can fully address this use case.

3\) Should this extension add a rotation capability to vkCmdResolveImage2KHR?

**RESOLVED** No. Use of vkCmdResolveImage2KHR is very slow and extremely bandwidth intensive on Qualcomm’s GPU architecture and use of pResolveAttachments in vkRenderPass is the strongly preferred approach. Therefore, we choose not to introduce a rotation capability to vkCmdResolveImage2KHR.

## [](#_new_structures)New Structures

- Extending [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2.html), [VkImageBlit2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2.html):
  
  - [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyCommandTransformInfoQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_ROTATED_COPY_COMMANDS_EXTENSION_NAME`
- `VK_QCOM_ROTATED_COPY_COMMANDS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COPY_COMMAND_TRANSFORM_INFO_QCOM`

## [](#_version_history)Version History

- Revision 1, 2020-09-19 (Jeff Leger)
- Revision 2, 2023-12-13 (Matthew Netsch)
  
  - Relax dependency on VK\_KHR\_swapchain

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_rotated_copy_commands)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0