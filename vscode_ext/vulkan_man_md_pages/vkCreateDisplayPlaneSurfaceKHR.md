# vkCreateDisplayPlaneSurfaceKHR(3) Manual Page

## Name

vkCreateDisplayPlaneSurfaceKHR - Create a slink:VkSurfaceKHR structure
representing a display plane and mode



## <a href="#_c_specification" class="anchor"></a>C Specification

A complete display configuration includes a mode, one or more display
planes and any parameters describing their behavior, and parameters
describing some aspects of the images associated with those planes.
Display surfaces describe the configuration of a single plane within a
complete display configuration. To create a `VkSurfaceKHR` object for a
display plane, call:

``` c
// Provided by VK_KHR_display
VkResult vkCreateDisplayPlaneSurfaceKHR(
    VkInstance                                  instance,
    const VkDisplaySurfaceCreateInfoKHR*        pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance corresponding to the physical device the
  targeted display is on.

- `pCreateInfo` is a pointer to a
  [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html)
  structure specifying which mode, plane, and other parameters to use,
  as described below.

- `pAllocator` is the allocator used for host memory allocated for the
  surface object when there is no more specific allocator available (see
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle
  in which the created surface is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDisplayPlaneSurfaceKHR-instance-parameter"
  id="VUID-vkCreateDisplayPlaneSurfaceKHR-instance-parameter"></a>
  VUID-vkCreateDisplayPlaneSurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateDisplayPlaneSurfaceKHR-pCreateInfo-parameter"
  id="VUID-vkCreateDisplayPlaneSurfaceKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateDisplayPlaneSurfaceKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateDisplayPlaneSurfaceKHR-pAllocator-parameter"
  id="VUID-vkCreateDisplayPlaneSurfaceKHR-pAllocator-parameter"></a>
  VUID-vkCreateDisplayPlaneSurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDisplayPlaneSurfaceKHR-pSurface-parameter"
  id="VUID-vkCreateDisplayPlaneSurfaceKHR-pSurface-parameter"></a>
  VUID-vkCreateDisplayPlaneSurfaceKHR-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDisplaySurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateInfoKHR.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDisplayPlaneSurfaceKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
