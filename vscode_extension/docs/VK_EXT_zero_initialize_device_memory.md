# VK\_EXT\_zero\_initialize\_device\_memory(3) Manual Page

## Name

VK\_EXT\_zero\_initialize\_device\_memory - device extension



## [](#_registered_extension_number)Registered Extension Number

621

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_zero_initialize_device_memory%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_zero_initialize_device_memory%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_zero\_initialize\_device\_memory](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_zero_initialize_device_memory.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-04-09

**Interactions and External Dependencies**

- Interacts with Vulkan 1.1.
- Interacts with `VK_KHR_get_physical_device_properties2`.

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, VALVE
- Mike Blumenkrantz, VALVE
- Tobias Hector, AMD
- Faith Ekstrand, Collabora
- Ricardo Garcia, Igalia
- Jan-Harald Fredriksen, ARM
- Spencer Fricke, LunarG

## [](#_description)Description

By default, Vulkan provides no guarantees that device memory allocated through vkAllocateMemory is cleared to zero. This means that applications wanting resources to be zero-initialized must execute a command such as vkCmdFillBuffer or vkCmdClearColorImage on the device to ensure a deterministic result. This can be wasteful if the underlying platform either:

- Already performs that zero clear anyway, due to e.g. security concerns.
- Can be performed more efficiently in implementation, by e.g. clearing pages to zero in the background after device memory is freed.

This extension also has uses in API layering and porting efforts, where zero memory behavior may be more strict than Vulkan. Different OS platforms also have wildly different behaviors here, which leads to implementations needing to apply workarounds to paper over these issues in the wild. If an extension exists to make allocation behavior explicit, we hopefully achieve a more robust ecosystem for Vulkan.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceZeroInitializeDeviceMemoryFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_ZERO_INITIALIZE_DEVICE_MEMORY_EXTENSION_NAME`
- `VK_EXT_ZERO_INITIALIZE_DEVICE_MEMORY_SPEC_VERSION`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`
- Extending [VkMemoryAllocateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagBits.html):
  
  - `VK_MEMORY_ALLOCATE_ZERO_INITIALIZE_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_DEVICE_MEMORY_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2025-03-10 (Mike Blumenkrantz)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_zero_initialize_device_memory).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0