# VK\_KHR\_copy\_commands2(3) Manual Page

## Name

VK\_KHR\_copy\_commands2 - device extension



## [](#_registered_extension_number)Registered Extension Number

338

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_copy_commands2%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_copy_commands2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-07-06

**Contributors**

- Jeff Leger, Qualcomm
- Tobias Hector, AMD
- Jan-Harald Fredriksen, ARM
- Tom Olson, ARM

## [](#_description)Description

This extension provides extensible versions of the Vulkan buffer and image copy commands. The new commands are functionally identical to the core commands, except that their copy parameters are specified using extensible structures that can be used to pass extension-specific information.

The following extensible copy commands are introduced with this extension: [vkCmdCopyBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBuffer2KHR.html), [vkCmdCopyImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage2KHR.html), [vkCmdCopyBufferToImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2KHR.html), [vkCmdCopyImageToBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2KHR.html), [vkCmdBlitImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2KHR.html), and [vkCmdResolveImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResolveImage2KHR.html). Each command contains an `*Info2KHR` structure parameter that includes `sType`/`pNext` members. Lower level structures describing each region to be copied are also extended with `sType`/`pNext` members.

## [](#_new_commands)New Commands

- [vkCmdBlitImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBlitImage2KHR.html)
- [vkCmdCopyBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBuffer2KHR.html)
- [vkCmdCopyBufferToImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyBufferToImage2KHR.html)
- [vkCmdCopyImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage2KHR.html)
- [vkCmdCopyImageToBuffer2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImageToBuffer2KHR.html)
- [vkCmdResolveImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResolveImage2KHR.html)

## [](#_new_structures)New Structures

- [VkBlitImageInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2KHR.html)
- [VkBufferCopy2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCopy2KHR.html)
- [VkBufferImageCopy2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferImageCopy2KHR.html)
- [VkCopyBufferInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferInfo2KHR.html)
- [VkCopyBufferToImageInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyBufferToImageInfo2KHR.html)
- [VkCopyImageInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2KHR.html)
- [VkCopyImageToBufferInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToBufferInfo2KHR.html)
- [VkImageBlit2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageBlit2KHR.html)
- [VkImageCopy2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2KHR.html)
- [VkImageResolve2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageResolve2KHR.html)
- [VkResolveImageInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveImageInfo2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_COPY_COMMANDS_2_EXTENSION_NAME`
- `VK_KHR_COPY_COMMANDS_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BLIT_IMAGE_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_BUFFER_COPY_2_KHR`
  - `VK_STRUCTURE_TYPE_BUFFER_IMAGE_COPY_2_KHR`
  - `VK_STRUCTURE_TYPE_COPY_BUFFER_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_COPY_BUFFER_TO_IMAGE_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_COPY_IMAGE_TO_BUFFER_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_BLIT_2_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_COPY_2_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_RESOLVE_2_KHR`
  - `VK_STRUCTURE_TYPE_RESOLVE_IMAGE_INFO_2_KHR`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2020-07-06 (Jeff Leger)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_copy_commands2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0