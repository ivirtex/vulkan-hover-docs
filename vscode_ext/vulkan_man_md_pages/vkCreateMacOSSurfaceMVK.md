# vkCreateMacOSSurfaceMVK(3) Manual Page

## Name

vkCreateMacOSSurfaceMVK - Create a VkSurfaceKHR object for a macOS
NSView



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a `VkSurfaceKHR` object for a macOS `NSView` or
[CAMetalLayer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/CAMetalLayer.html), call:

``` c
// Provided by VK_MVK_macos_surface
VkResult vkCreateMacOSSurfaceMVK(
    VkInstance                                  instance,
    const VkMacOSSurfaceCreateInfoMVK*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance with which to associate the surface.

- `pCreateInfo` is a pointer to a
  [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateInfoMVK.html)
  structure containing parameters affecting the creation of the surface
  object.

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
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>vkCreateMacOSSurfaceMVK</code> function is considered
deprecated and has been superseded by <a
href="vkCreateMetalSurfaceEXT.html">vkCreateMetalSurfaceEXT</a> from the
<a
href="VK_EXT_metal_surface.html"><code>VK_EXT_metal_surface</code></a>
extension.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkCreateMacOSSurfaceMVK-instance-parameter"
  id="VUID-vkCreateMacOSSurfaceMVK-instance-parameter"></a>
  VUID-vkCreateMacOSSurfaceMVK-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateMacOSSurfaceMVK-pCreateInfo-parameter"
  id="VUID-vkCreateMacOSSurfaceMVK-pCreateInfo-parameter"></a>
  VUID-vkCreateMacOSSurfaceMVK-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateInfoMVK.html)
  structure

- <a href="#VUID-vkCreateMacOSSurfaceMVK-pAllocator-parameter"
  id="VUID-vkCreateMacOSSurfaceMVK-pAllocator-parameter"></a>
  VUID-vkCreateMacOSSurfaceMVK-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateMacOSSurfaceMVK-pSurface-parameter"
  id="VUID-vkCreateMacOSSurfaceMVK-pSurface-parameter"></a>
  VUID-vkCreateMacOSSurfaceMVK-pSurface-parameter  
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

[VK_MVK_macos_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MVK_macos_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html),
[VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateInfoMVK.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateMacOSSurfaceMVK"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
