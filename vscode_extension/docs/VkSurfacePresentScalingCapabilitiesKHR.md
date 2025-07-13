# VkSurfacePresentScalingCapabilitiesKHR(3) Manual Page

## Name

VkSurfacePresentScalingCapabilitiesKHR - Structure describing the presentation scaling capabilities of the surface



## [](#_c_specification)C Specification

The `VkSurfacePresentScalingCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_surface_maintenance1
typedef struct VkSurfacePresentScalingCapabilitiesKHR {
    VkStructureType             sType;
    void*                       pNext;
    VkPresentScalingFlagsKHR    supportedPresentScaling;
    VkPresentGravityFlagsKHR    supportedPresentGravityX;
    VkPresentGravityFlagsKHR    supportedPresentGravityY;
    VkExtent2D                  minScaledImageExtent;
    VkExtent2D                  maxScaledImageExtent;
} VkSurfacePresentScalingCapabilitiesKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_surface_maintenance1
typedef VkSurfacePresentScalingCapabilitiesKHR VkSurfacePresentScalingCapabilitiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `supportedPresentScaling` is a bitmask of [VkPresentScalingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagBitsKHR.html) representing the scaling methods supported by the surface, or `0` if application-defined scaling is not supported.
- `supportedPresentGravityX` is a bitmask of [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html) representing the X-axis pixel gravity supported by the surface, or `0` if Vulkan-defined pixel gravity is not supported for the X axis.
- `supportedPresentGravityY` is a bitmask of [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html) representing the Y-axis pixel gravity supported by the surface, or `0` if Vulkan-defined pixel gravity is not supported for the Y axis.
- `minScaledImageExtent` contains the smallest valid swapchain extent for the surface on the specified device when one of the scaling methods specified in `supportedPresentScaling` is used, or the special value (0xFFFFFFFF, 0xFFFFFFFF) indicating that the surface size will be determined by the extent of a swapchain targeting the surface. The `width` and `height` of the extent will each be smaller than or equal to the corresponding `width` and `height` of [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`minImageExtent`.
- `maxScaledImageExtent` contains the largest valid swapchain extent for the surface on the specified device when one of the scaling methods specified in `supportedPresentScaling` is used, or the special value described above for `minScaledImageExtent`. The `width` and `height` of the extent will each be greater than or equal to the corresponding `width` and `height` of [VkSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesKHR.html)::`maxImageExtent`.

## [](#_description)Description

To query the set of supported scaling modes for a given present mode, add a [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html) structure in the `pNext` chain of [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) when calling [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html). The implementation **must** return the same values in `VkSurfacePresentScalingCapabilitiesKHR` for any of the compatible present modes as obtained through [VkSurfacePresentModeCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeCompatibilityKHR.html).

The application **can** specify the scaling mode when creating a swapchain through the use of [VkSwapchainPresentScalingCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentScalingCreateInfoKHR.html).

Valid Usage (Implicit)

- [](#VUID-VkSurfacePresentScalingCapabilitiesKHR-sType-sType)VUID-VkSurfacePresentScalingCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_PRESENT_SCALING_CAPABILITIES_KHR`
- [](#VUID-VkSurfacePresentScalingCapabilitiesKHR-supportedPresentScaling-parameter)VUID-VkSurfacePresentScalingCapabilitiesKHR-supportedPresentScaling-parameter  
  `supportedPresentScaling` **must** be a valid combination of [VkPresentScalingFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagBitsKHR.html) values
- [](#VUID-VkSurfacePresentScalingCapabilitiesKHR-supportedPresentGravityX-parameter)VUID-VkSurfacePresentScalingCapabilitiesKHR-supportedPresentGravityX-parameter  
  `supportedPresentGravityX` **must** be a valid combination of [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html) values
- [](#VUID-VkSurfacePresentScalingCapabilitiesKHR-supportedPresentGravityY-parameter)VUID-VkSurfacePresentScalingCapabilitiesKHR-supportedPresentGravityY-parameter  
  `supportedPresentGravityY` **must** be a valid combination of [VkPresentGravityFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagBitsKHR.html) values

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VK\_KHR\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_maintenance1.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkPresentGravityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentGravityFlagsKHR.html), [VkPresentScalingFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentScalingFlagsKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfacePresentScalingCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0