# VK\_KHR\_present\_id2(3) Manual Page

## Name

VK\_KHR\_present\_id2 - device extension



## [](#_registered_extension_number)Registered Extension Number

480

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)  
and  
[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Daniel Stone

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_present\_id2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_present_id2.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-06

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, Valve
- James Jones, NVIDIA
- Daniel Stone, Collabora
- Derek Foreman, Collabora
- *contributors to \`[VK\_KHR\_present\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id.html)\`*

## [](#_description)Description

This device extension allows an application that uses the `VK_KHR_swapchain` extension to provide an identifier for present operations on a swapchain. An application **can** use this to reference specific present operations in other extensions.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePresentId2FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePresentId2FeaturesKHR.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfaceCapabilitiesPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentId2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PRESENT_ID_2_EXTENSION_NAME`
- `VK_KHR_PRESENT_ID_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_ID_2_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PRESENT_ID_2_KHR`
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_ID_2_KHR`
- Extending [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html):
  
  - `VK_SWAPCHAIN_CREATE_PRESENT_ID_2_BIT_KHR`

## [](#_issues)Issues

None.

## [](#_examples)Examples

## [](#_version_history)Version History

- Revision 1, 2022-05-10 (Daniel Stone)
  
  - Repurposed VK\_KHR\_present\_id to be driven by surface capabilities

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_present_id2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0