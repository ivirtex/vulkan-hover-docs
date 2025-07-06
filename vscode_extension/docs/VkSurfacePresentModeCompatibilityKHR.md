# VkSurfacePresentModeCompatibilityKHR(3) Manual Page

## Name

VkSurfacePresentModeCompatibilityKHR - Structure describing the subset of compatible presentation modes for the purposes of switching without swapchain recreation



## [](#_c_specification)C Specification

The `VkSurfacePresentModeCompatibilityKHR` structure is defined as:

```c++
// Provided by VK_KHR_surface_maintenance1
typedef struct VkSurfacePresentModeCompatibilityKHR {
    VkStructureType      sType;
    void*                pNext;
    uint32_t             presentModeCount;
    VkPresentModeKHR*    pPresentModes;
} VkSurfacePresentModeCompatibilityKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_surface_maintenance1
typedef VkSurfacePresentModeCompatibilityKHR VkSurfacePresentModeCompatibilityEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentModeCount` is an integer related to the number of present modes available or queried, as described below.
- `pPresentModes` is a pointer to an array of [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) in which present modes compatible with a given present mode are returned.

## [](#_description)Description

If `pPresentModes` is `NULL`, then the number of present modes that are compatible with the one specified in [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html) is returned in `presentModeCount`. Otherwise, `presentModeCount` **must** point to a variable set by the application to the number of elements in the `pPresentModes` array, and on return the variable is overwritten with the number of values actually written to `pPresentModes`. If the value of `presentModeCount` is less than the number of compatible present modes that are supported, at most `presentModeCount` values will be written to `pPresentModes`. The implementation **must** include the present mode passed to [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html) in `pPresentModes`, unless `presentModeCount` is zero.

To query the set of present modes compatible with a given initial present mode, add a [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html) structure in the `pNext` chain of [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html) when calling [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).

The application **can** create a swapchain whose present mode **can** be modified through the use of [VkSwapchainPresentModesCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentModesCreateInfoKHR.html).

Valid Usage (Implicit)

- [](#VUID-VkSurfacePresentModeCompatibilityKHR-sType-sType)VUID-VkSurfacePresentModeCompatibilityKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_COMPATIBILITY_KHR`
- [](#VUID-VkSurfacePresentModeCompatibilityKHR-pPresentModes-parameter)VUID-VkSurfacePresentModeCompatibilityKHR-pPresentModes-parameter  
  If `presentModeCount` is not `0`, and `pPresentModes` is not `NULL`, `pPresentModes` **must** be a valid pointer to an array of `presentModeCount` [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html) values

## [](#_see_also)See Also

[VK\_EXT\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_surface_maintenance1.html), [VK\_KHR\_surface\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_maintenance1.html), [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfacePresentModeCompatibilityKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0