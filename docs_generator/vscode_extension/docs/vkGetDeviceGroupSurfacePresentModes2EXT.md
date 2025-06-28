# vkGetDeviceGroupSurfacePresentModes2EXT(3) Manual Page

## Name

vkGetDeviceGroupSurfacePresentModes2EXT - Query device group present capabilities for a surface



## [](#_c_specification)C Specification

Alternatively, to query the supported device group presentation modes for a surface combined with select other fixed swapchain creation parameters, call:

```c++
// Provided by VK_EXT_full_screen_exclusive with VK_KHR_device_group or VK_VERSION_1_1
VkResult vkGetDeviceGroupSurfacePresentModes2EXT(
    VkDevice                                    device,
    const VkPhysicalDeviceSurfaceInfo2KHR*      pSurfaceInfo,
    VkDeviceGroupPresentModeFlagsKHR*           pModes);
```

## [](#_parameters)Parameters

- `device` is the logical device.
- `pSurfaceInfo` is a pointer to a [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure describing the surface and other fixed parameters that would be consumed by [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSwapchainKHR.html).
- `pModes` is a pointer to a [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagsKHR.html) in which the supported device group present modes for the surface are returned.

## [](#_description)Description

`vkGetDeviceGroupSurfacePresentModes2EXT` behaves similarly to [vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html), with the ability to specify extended inputs via chained input structures.

Valid Usage

- [](#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-06213)VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-06213  
  `pSurfaceInfo->surface` **must** be supported by all physical devices associated with `device`, as reported by [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html) or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-device-parameter)VUID-vkGetDeviceGroupSurfacePresentModes2EXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-parameter)VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pSurfaceInfo-parameter  
  `pSurfaceInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) structure
- [](#VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pModes-parameter)VUID-vkGetDeviceGroupSurfacePresentModes2EXT-pModes-parameter  
  `pModes` **must** be a valid pointer to a [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagsKHR.html) value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_SURFACE_LOST_KHR`

## [](#_see_also)See Also

[VK\_EXT\_full\_screen\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_full_screen_exclusive.html), [VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupPresentModeFlagsKHR.html), [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceGroupSurfacePresentModes2EXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0