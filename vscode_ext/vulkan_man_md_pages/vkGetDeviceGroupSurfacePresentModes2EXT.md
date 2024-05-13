# vkGetDeviceGroupSurfacePresentModes2EXT(3) Manual Page

## Name

vkGetDeviceGroupSurfacePresentModes2EXT - Query device group present
capabilities for a surface



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to query the supported device group presentation modes
for a surface combined with select other fixed swapchain creation
parameters, call:

``` c
// Provided by VK_VERSION_1_1 with VK_EXT_full_screen_exclusive, VK_KHR_device_group with VK_EXT_full_screen_exclusive
VkResult vkGetDeviceGroupSurfacePresentModes2EXT(
    VkDevice                                    device,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    VkDeviceGroupPresentModeFlagsKHR*           pModes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device.

- `pSurfaceInfo` is a pointer to a
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure describing the surface and other fixed parameters that would
  be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html).

- `pModes` is a pointer to a
  [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html)
  in which the supported device group present modes for the surface are
  returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetDeviceGroupSurfacePresentModes2EXT` behaves similarly to
[vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html),
with the ability to specify extended inputs via chained input
structures.

Valid Usage

- <a
  href="#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-06213"
  id="VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-06213"></a>
  VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-06213  
  `pSurfaceInfo->surface` **must** be supported by all physical devices
  associated with `device`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-device-parameter"
  id="VUID-vkGetDeviceGroupSurfacePresentModes2EXT-device-parameter"></a>
  VUID-vkGetDeviceGroupSurfacePresentModes2EXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-parameter"
  id="VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-parameter"></a>
  VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
  structure

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pModes-parameter"
  id="VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pModes-parameter"></a>
  VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pModes-parameter  
  `pModes` **must** be a valid pointer to a
  [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html)
  value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_full_screen_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_full_screen_exclusive.html),
[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html),
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceGroupSurfacePresentModes2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
