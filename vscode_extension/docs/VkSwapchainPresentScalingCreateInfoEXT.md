# VkSwapchainPresentScalingCreateInfoEXT(3) Manual Page

## Name

VkSwapchainPresentScalingCreateInfoEXT - Scaling behavior when presenting to the surface



## [](#_c_specification)C Specification

When an application presents a swapchain image with dimensions different than those of the target surface, different behavior is possible on different platforms per their respective specifications:

- Presentation fails and `VK_ERROR_OUT_OF_DATE_KHR` is returned
- Scaling is done and `VK_SUCCESS` or `VK_SUBOPTIMAL_KHR` is returned
- Unspecified scaling using an arbitrary combination of stretching, centering and/or clipping.

Applications **can** define specific behavior when creating a swapchain by including the `VkSwapchainPresentScalingCreateInfoEXT` structure in the `pNext` chain of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure.

The `VkSwapchainPresentScalingCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentScalingCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkPresentScalingFlagsEXT    scalingBehavior;
    VkPresentGravityFlagsEXT    presentGravityX;
    VkPresentGravityFlagsEXT    presentGravityY;
} VkSwapchainPresentScalingCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `scalingBehavior` is `0` or the scaling method to use when the dimensions of the surface and swapchain images differ.
- `presentGravityX` is `0` or the x-axis direction in which swapchain image pixels gravitate relative to the surface when `scalingBehavior` does not result in a one-to-one pixel mapping between the scaled swapchain image and the surface.
- `presentGravityY` is `0` or the y-axis direction in which swapchain image pixels gravitate relative to the surface when `scalingBehavior` does not result in a one-to-one pixel mapping between the scaled swapchain image and the surface.

## [](#_description)Description

If `scalingBehavior` is `0`, the result of presenting a swapchain image with dimensions that do not match the surface dimensions is implementation and platform-dependent. If `presentGravityX` or `presentGravityY` are `0`, the presentation gravity **must** match that defined by the native platform surface on platforms which define surface gravity.

Valid Usage

- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07765)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07765  
  If `presentGravityX` is `0`, `presentGravityY` **must** be `0`
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07766)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07766  
  If `presentGravityX` is not `0`, `presentGravityY` **must** not be `0`
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07767)VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07767  
  `scalingBehavior` **must** not have more than one bit set
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07768)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07768  
  `presentGravityX` **must** not have more than one bit set
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07769)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07769  
  `presentGravityY` **must** not have more than one bit set
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07770)VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07770  
  `scalingBehavior` **must** be `0` or a valid scaling method for the surface as returned in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentScaling`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07771)VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07771  
  If the swapchain is created with [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html), `scalingBehavior` **must** be `0` or a valid scaling method for the surface as returned in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentScaling`, given each present mode in [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07772)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07772  
  `presentGravityX` **must** be `0` or a valid x-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityX`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07773)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07773  
  If the swapchain is created with [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html), `presentGravityX` **must** be `0` or a valid x-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityX`, given each present mode in [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07774)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07774  
  `presentGravityY` **must** be `0` or a valid y-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityY`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07775)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07775  
  If the swapchain is created with [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html), `presentGravityY` **must** be `0` or a valid y-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityY`, given each present mode in [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes` in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeEXT.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-swapchainMaintenance1-10154)VUID-VkSwapchainPresentScalingCreateInfoEXT-swapchainMaintenance1-10154  
  If the [`swapchainMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-swapchainMaintenance1) feature is not enabled, then `scalingBehavior`, `presentGravityX`, and `presentGravityY` **must** be `0`

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-sType-sType)VUID-VkSwapchainPresentScalingCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_SCALING_CREATE_INFO_EXT`
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-parameter)VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-parameter  
  `scalingBehavior` **must** be a valid combination of [VkPresentScalingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagBitsEXT.html) values
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-parameter)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-parameter  
  `presentGravityX` **must** be a valid combination of [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsEXT.html) values
- [](#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-parameter)VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-parameter  
  `presentGravityY` **must** be a valid combination of [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsEXT.html) values

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagsEXT.html), [VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentScalingCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0