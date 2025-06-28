# VK\_KHR\_bind\_memory2(3) Manual Page

## Name

VK\_KHR\_bind\_memory2 - device extension



## [](#_registered_extension_number)Registered Extension Number

158

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_bind_memory2%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_bind_memory2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Tobias Hector, Imagination Technologies

## [](#_description)Description

This extension provides versions of [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory.html) and [vkBindImageMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindImageMemory.html) that allow multiple bindings to be performed at once, and are extensible.

This extension also introduces `VK_IMAGE_CREATE_ALIAS_BIT_KHR`, which allows “identical” images that alias the same memory to interpret the contents consistently, even across image layout changes.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkBindBufferMemory2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory2KHR.html)
- [vkBindImageMemory2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindImageMemory2KHR.html)

## [](#_new_structures)New Structures

- [VkBindBufferMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfoKHR.html)
- [VkBindImageMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_BIND_MEMORY_2_EXTENSION_NAME`
- `VK_KHR_BIND_MEMORY_2_SPEC_VERSION`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_ALIAS_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO_KHR`
  - `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR`

## [](#_version_history)Version History

- Revision 1, 2017-05-19 (Tobias Hector)
  
  - Pulled bind memory functions into their own extension

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_bind_memory2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0