# VK_NV_present_barrier(3) Manual Page

## Name

VK_NV_present_barrier - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

293

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  
and  
[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  
and  
[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Liya Li <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_present_barrier%5D%20@liyli%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_present_barrier%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>liyli</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds support for synchronizing corresponding presentation
requests across multiple swapchains using the *present barrier*.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePresentBarrierFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePresentBarrierFeaturesNV.html)

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkSurfaceCapabilitiesPresentBarrierNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilitiesPresentBarrierNV.html)

- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentBarrierCreateInfoNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_PRESENT_BARRIER_EXTENSION_NAME`

- `VK_NV_PRESENT_BARRIER_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_BARRIER_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_BARRIER_NV`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_BARRIER_CREATE_INFO_NV`

## <a href="#_issues" class="anchor"></a>Issues

1\) Is there a query interface to check if a swapchain is using the
present barrier?

**RESOLVED**. There is no such query interface. When creating a
swapchain, an application can specify to use the *present barrier*, and
if the swapchain is created successfully, this swapchain will be using
the present barrier.

2\) Do we need an extra interface to set up the present barrier across
distributed systems?

**RESOLVED**. If the required hardware is presented in the system, and
all settings for the physical synchronization with other systems are set
up, an implementation manages the configuration automatically when
creating a swapchain, without any extra calls from the application.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-07-20

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_present_barrier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
