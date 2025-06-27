# vkCreateIOSSurfaceMVK(3) Manual Page

## Name

vkCreateIOSSurfaceMVK - Create a VkSurfaceKHR object for an iOS UIView



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a `VkSurfaceKHR` object for an iOS `UIView` or
[CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html), call:

``` c
// Provided by VK_MVK_ios_surface
VkResult vkCreateIOSSurfaceMVK(
    VkInstance                                  instance,
    const VkIOSSurfaceCreateInfoMVK*            pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance with which to associate the surface.

- `pCreateInfo` is a pointer to a
  [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateInfoMVK.html) structure
  containing parameters affecting the creation of the surface object.

- `pAllocator` is the allocator used for host memory allocated for the
  surface object when there is no more specific allocator available (see
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle
  in which the created surface object is returned.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>vkCreateIOSSurfaceMVK</code> function is considered
deprecated and has been superseded by <a
href="vkCreateMetalSurfaceEXT.html">vkCreateMetalSurfaceEXT</a> from the
<a
href="VK_EXT_metal_surface.html"><code>VK_EXT_metal_surface</code></a>
extension.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkCreateIOSSurfaceMVK-instance-parameter"
  id="VUID-vkCreateIOSSurfaceMVK-instance-parameter"></a>
  VUID-vkCreateIOSSurfaceMVK-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateIOSSurfaceMVK-pCreateInfo-parameter"
  id="VUID-vkCreateIOSSurfaceMVK-pCreateInfo-parameter"></a>
  VUID-vkCreateIOSSurfaceMVK-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateInfoMVK.html) structure

- <a href="#VUID-vkCreateIOSSurfaceMVK-pAllocator-parameter"
  id="VUID-vkCreateIOSSurfaceMVK-pAllocator-parameter"></a>
  VUID-vkCreateIOSSurfaceMVK-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateIOSSurfaceMVK-pSurface-parameter"
  id="VUID-vkCreateIOSSurfaceMVK-pSurface-parameter"></a>
  VUID-vkCreateIOSSurfaceMVK-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_NATIVE_WINDOW_IN_USE_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MVK_ios_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MVK_ios_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkIOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateInfoMVK.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateIOSSurfaceMVK"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
