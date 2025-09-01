# VkSwapchainPresentScalingCreateInfoKHR(3) Manual Page

## Name

VkSwapchainPresentScalingCreateInfoKHR - Scaling behavior when presenting to the surface



## [](#_c_specification)C Specification

When an application presents a swapchain image with dimensions different than those of the target surface, different behavior is possible on different platforms per their respective specifications:

- Presentation fails and `VK_ERROR_OUT_OF_DATE_KHR` is returned
- Scaling is done and `VK_SUCCESS` or `VK_SUBOPTIMAL_KHR` is returned
- Unspecified scaling using an arbitrary combination of stretching, centering and/or clipping.

Applications **can** define specific behavior when creating a swapchain by including the `VkSwapchainPresentScalingCreateInfoKHR` structure in the `pNext` chain of the [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) structure.

The `VkSwapchainPresentScalingCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_swapchain_maintenance1
typedef struct VkSwapchainPresentScalingCreateInfoKHR {
    VkStructureType             sType;
    const void*                 pNext;
    VkPresentScalingFlagsKHR    scalingBehavior;
    VkPresentGravityFlagsKHR    presentGravityX;
    VkPresentGravityFlagsKHR    presentGravityY;
} VkSwapchainPresentScalingCreateInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef VkSwapchainPresentScalingCreateInfoKHR VkSwapchainPresentScalingCreateInfoEXT;
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

- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07765)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07765  
  If `presentGravityX` is `0`, `presentGravityY` **must** be `0`
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07766)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07766  
  If `presentGravityX` is not `0`, `presentGravityY` **must** not be `0`
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-07767)VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-07767  
  `scalingBehavior` **must** not have more than one bit set
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07768)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07768  
  `presentGravityX` **must** not have more than one bit set
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-07769)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-07769  
  `presentGravityY` **must** not have more than one bit set
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-07770)VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-07770  
  `scalingBehavior` **must** be `0` or a valid scaling method for the surface as returned in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentScaling`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-07771)VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-07771  
  If the swapchain is created with [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html), `scalingBehavior` **must** be `0` or a valid scaling method for the surface as returned in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentScaling`, given each present mode in [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html)::`pPresentModes` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07772)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07772  
  `presentGravityX` **must** be `0` or a valid x-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentGravityX`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07773)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-07773  
  If the swapchain is created with [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html), `presentGravityX` **must** be `0` or a valid x-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentGravityX`, given each present mode in [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html)::`pPresentModes` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-07774)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-07774  
  `presentGravityY` **must** be `0` or a valid y-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentGravityY`, given [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-07775)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-07775  
  If the swapchain is created with [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html), `presentGravityY` **must** be `0` or a valid y-axis present gravity for the surface as returned in [VkSurfacePresentScalingCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentScalingCapabilitiesKHR.html)::`supportedPresentGravityY`, given each present mode in [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html)::`pPresentModes` in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-swapchainMaintenance1-10154)VUID-VkSwapchainPresentScalingCreateInfoKHR-swapchainMaintenance1-10154  
  If the [`swapchainMaintenance1`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-swapchainMaintenance1) feature is not enabled, then `scalingBehavior`, `presentGravityX`, and `presentGravityY` **must** be `0`

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-sType-sType)VUID-VkSwapchainPresentScalingCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_SCALING_CREATE_INFO_KHR`
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-parameter)VUID-VkSwapchainPresentScalingCreateInfoKHR-scalingBehavior-parameter  
  `scalingBehavior` **must** be a valid combination of [VkPresentScalingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagBitsKHR.html) values
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-parameter)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityX-parameter  
  `presentGravityX` **must** be a valid combination of [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html) values
- [](#VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-parameter)VUID-VkSwapchainPresentScalingCreateInfoKHR-presentGravityY-parameter  
  `presentGravityY` **must** be a valid combination of [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VK\_KHR\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain_maintenance1.html), [VkPresentGravityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagsKHR.html), [VkPresentScalingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentScalingCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0