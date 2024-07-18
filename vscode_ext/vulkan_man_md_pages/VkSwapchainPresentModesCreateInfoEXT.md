# VkSwapchainPresentModesCreateInfoEXT(3) Manual Page

## Name

VkSwapchainPresentModesCreateInfoEXT - All presentation modes usable by
the swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

Applications **can** modify the presentation mode used by the swapchain
on a per-presentation basis. However, all presentation modes the
application intends to use with the swapchain **must** be specified at
swapchain creation time. To specify more than one presentation mode when
creating a swapchain, include the `VkSwapchainPresentModesCreateInfoEXT`
structure in the `pNext` chain of the
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure.

The `VkSwapchainPresentModesCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentModesCreateInfoEXT {
    VkStructureType            sType;
    const void*                pNext;
    uint32_t                   presentModeCount;
    const VkPresentModeKHR*    pPresentModes;
} VkSwapchainPresentModesCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentModeCount` is the number of presentation modes provided.

- `pPresentModes` is a list of presentation modes with
  `presentModeCount` entries

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSwapchainPresentModesCreateInfoEXT-None-07762"
  id="VUID-VkSwapchainPresentModesCreateInfoEXT-None-07762"></a>
  VUID-VkSwapchainPresentModesCreateInfoEXT-None-07762  
  Each entry in pPresentModes **must** be one of the
  [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) values returned by
  `vkGetPhysicalDeviceSurfacePresentModesKHR` for the surface

- <a href="#VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-07763"
  id="VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-07763"></a>
  VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-07763  
  The entries in pPresentModes **must** be a subset of the present modes
  returned in
  [VkSurfacePresentModeCompatibilityEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeCompatibilityEXT.html)::`pPresentModes`,
  given
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a href="#VUID-VkSwapchainPresentModesCreateInfoEXT-presentMode-07764"
  id="VUID-VkSwapchainPresentModesCreateInfoEXT-presentMode-07764"></a>
  VUID-VkSwapchainPresentModesCreateInfoEXT-presentMode-07764  
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode`
  **must** be included in `pPresentModes`

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainPresentModesCreateInfoEXT-sType-sType"
  id="VUID-VkSwapchainPresentModesCreateInfoEXT-sType-sType"></a>
  VUID-VkSwapchainPresentModesCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_MODES_CREATE_INFO_EXT`

- <a
  href="#VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-parameter"
  id="VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-parameter"></a>
  VUID-VkSwapchainPresentModesCreateInfoEXT-pPresentModes-parameter  
  `pPresentModes` **must** be a valid pointer to an array of
  `presentModeCount` valid [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html)
  values

- <a
  href="#VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeCount-arraylength"
  id="VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeCount-arraylength"></a>
  VUID-VkSwapchainPresentModesCreateInfoEXT-presentModeCount-arraylength  
  `presentModeCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainPresentModesCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
