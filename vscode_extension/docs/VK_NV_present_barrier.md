# VK\_NV\_present\_barrier(3) Manual Page

## Name

VK\_NV\_present\_barrier - device extension



## [](#_registered_extension_number)Registered Extension Number

293

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)  
and  
[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Liya Li [\[GitHub\]liyli](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_present_barrier%5D%20%40liyli%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_present_barrier%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-05-16

**Contributors**

- Liya Li, Nvidia
- Martin Schwarzer, Nvidia
- Andy Wolf, Nvidia
- Ian Williams, Nvidia
- Ben Morris, Nvidia
- James Jones, Nvidia
- Jeff Juliano, Nvidia

## [](#_description)Description

This extension adds support for synchronizing corresponding presentation requests across multiple swapchains using the *present barrier*.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePresentBarrierFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePresentBarrierFeaturesNV.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfaceCapabilitiesPresentBarrierNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentBarrierNV.html)
- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentBarrierCreateInfoNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_PRESENT_BARRIER_EXTENSION_NAME`
- `VK_NV_PRESENT_BARRIER_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_BARRIER_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_BARRIER_NV`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_BARRIER_CREATE_INFO_NV`

## [](#_issues)Issues

1\) Is there a query interface to check if a swapchain is using the present barrier?

**RESOLVED**. There is no such query interface. When creating a swapchain, an application can specify to use the *present barrier*, and if the swapchain is created successfully, this swapchain will be using the present barrier.

2\) Do we need an extra interface to set up the present barrier across distributed systems?

**RESOLVED**. If the required hardware is presented in the system, and all settings for the physical synchronization with other systems are set up, an implementation manages the configuration automatically when creating a swapchain, without any extra calls from the application.

## [](#_version_history)Version History

- Revision 1, 2022-07-20
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_present_barrier)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0