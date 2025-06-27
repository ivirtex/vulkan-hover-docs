# VkSurfacePresentModeCompatibilityEXT(3) Manual Page

## Name

VkSurfacePresentModeCompatibilityEXT - Structure describing the subset
of compatible presentation modes for the purposes of switching without
swapchain recreation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfacePresentModeCompatibilityEXT` structure is defined as:

``` c
// Provided by VK_EXT_surface_maintenance1
typedef struct VkSurfacePresentModeCompatibilityEXT {
    VkStructureType      sType;
    void*                pNext;
    uint32_t             presentModeCount;
    VkPresentModeKHR*    pPresentModes;
} VkSurfacePresentModeCompatibilityEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentModeCount` is an integer related to the number of present
  modes available or queried, as described below.

- `pPresentModes` is a pointer to an array of
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) in which present modes
  compatible with a given present mode are returned.

## <a href="#_description" class="anchor"></a>Description

If `pPresentModes` is `NULL`, then the number of present modes that are
compatible with the one specified in
[VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html) is returned in
`presentModeCount`. Otherwise, `presentModeCount` **must** point to a
variable set by the application to the number of elements in the
`pPresentModes` array, and on return the variable is overwritten with
the number of values actually written to `pPresentModes`. If the value
of `presentModeCount` is less than the number of compatible present
modes that are supported, at most `presentModeCount` values will be
written to `pPresentModes`. The implementation **must** include the
present mode passed to
[VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html) in
`pPresentModes`, unless `presentModeCount` is zero.

Before creating a swapchain whose present modes **can** be modified
through the use of
[VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html),
obtain the set of present modes compatible with a given initial present
mode by including a
[VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html) structure in the
`pNext` chain of
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)
when calling
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html).

Valid Usage (Implicit)

- <a href="#VUID-VkSurfacePresentModeCompatibilityEXT-sType-sType"
  id="VUID-VkSurfacePresentModeCompatibilityEXT-sType-sType"></a>
  VUID-VkSurfacePresentModeCompatibilityEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_PRESENT_MODE_COMPATIBILITY_EXT`

- <a
  href="#VUID-VkSurfacePresentModeCompatibilityEXT-pPresentModes-parameter"
  id="VUID-VkSurfacePresentModeCompatibilityEXT-pPresentModes-parameter"></a>
  VUID-VkSurfacePresentModeCompatibilityEXT-pPresentModes-parameter  
  If `presentModeCount` is not `0`, and `pPresentModes` is not `NULL`,
  `pPresentModes` **must** be a valid pointer to an array of
  `presentModeCount` [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_surface_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_surface_maintenance1.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfacePresentModeCompatibilityEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
