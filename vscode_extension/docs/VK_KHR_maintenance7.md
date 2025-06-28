# VK\_KHR\_maintenance7(3) Manual Page

## Name

VK\_KHR\_maintenance7 - device extension



## [](#_registered_extension_number)Registered Extension Number

563

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance7%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance7%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_maintenance7](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_maintenance7.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-01-30

**Interactions and External Dependencies**

**Contributors**

- Mike Blumenkrantz, Valve
- Hans-Kristian Arntzen, Valve
- Pan Gao, Huawei
- Tobias Hector, AMD
- Jon Leech, Khronos
- Daniel Story, Nintendo
- Shahbaz Youssefi, Google
- Yiwei Zhang, Google
- Matthew Netsch, Qualcomm

## [](#_description)Description

[VK\_KHR\_maintenance7](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance7.html) adds a collection of minor features, none of which would warrant an entire extension of their own.

The proposed new features are as follows:

- Add a property query to determine if a framebuffer writes to depth or stencil aspect does not trigger a write access in the sibling aspect. For example, this allows sampling stencil aspect as a texture while rendering to the sibling depth attachment and vice-versa given appropriate image layouts.
- Add a way to query information regarding the underlying devices in environments where the Vulkan implementation is provided through layered implementations. For example, running on Mesa/Venus, driver ID is returned as `VK_DRIVER_ID_MESA_VENUS`, but it can be necessary to know what the real driver under the hood is. The new [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html) structure can be used to gather information regarding layers underneath the top-level physical device.
- Promote `VK_RENDERING_CONTENTS_INLINE_BIT_EXT` and `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_EXT` to KHR
- Add a limit to report the maximum total count of dynamic uniform buffers and dynamic storage buffers that can be included in a pipeline layout.
- Require that for an unsigned integer query, the 32-bit result value **must** be equal to the 32 least significant bits of the equivalent 64-bit result value.
- Add query for robust access support when using fragment shading rate attachments

## [](#_new_structures)New Structures

- [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMaintenance7FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance7FeaturesKHR.html)
- Extending [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html):
  
  - [VkPhysicalDeviceLayeredApiVulkanPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiVulkanPropertiesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceLayeredApiPropertiesListKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesListKHR.html)
  - [VkPhysicalDeviceMaintenance7PropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance7PropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkPhysicalDeviceLayeredApiKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE_7_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_7_SPEC_VERSION`
- Extending [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html):
  
  - `VK_RENDERING_CONTENTS_INLINE_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_PROPERTIES_LIST_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_API_VULKAN_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_7_PROPERTIES_KHR`
- Extending [VkSubpassContents](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassContents.html):
  
  - `VK_SUBPASS_CONTENTS_INLINE_AND_SECONDARY_COMMAND_BUFFERS_KHR`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2024-01-30 (Jon Leech)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance7)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0