# vkCreateAndroidSurfaceKHR(3) Manual Page

## Name

vkCreateAndroidSurfaceKHR - Create a slink:VkSurfaceKHR object for an
Android native window



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a `VkSurfaceKHR` object for an Android native window, call:

``` c
// Provided by VK_KHR_android_surface
VkResult vkCreateAndroidSurfaceKHR(
    VkInstance                                  instance,
    const VkAndroidSurfaceCreateInfoKHR*        pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance to associate the surface with.

- `pCreateInfo` is a pointer to a `VkAndroidSurfaceCreateInfoKHR`
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

During the lifetime of a surface created using a particular
[ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html) handle any attempts to create
another surface for the same [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html) and any
attempts to connect to the same [ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html)
through other platform mechanisms will fail.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>In particular, only one <code>VkSurfaceKHR</code>
<strong>can</strong> exist at a time for a given window. Similarly, a
native window <strong>cannot</strong> be used by both a
<code>VkSurfaceKHR</code> and <code>EGLSurface</code>
simultaneously.</p></td>
</tr>
</tbody>
</table>

If successful, `vkCreateAndroidSurfaceKHR` increments the
[ANativeWindow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/ANativeWindow.html)’s reference count, and
`vkDestroySurfaceKHR` will decrement it.

On Android, when a swapchain’s `imageExtent` does not match the
surface’s `currentExtent`, the presentable images will be scaled to the
surface’s dimensions during presentation. `minImageExtent` is (1,1), and
`maxImageExtent` is the maximum image size supported by the consumer.
For the system compositor, `currentExtent` is the window size (i.e. the
consumer’s preferred size).

Valid Usage (Implicit)

- <a href="#VUID-vkCreateAndroidSurfaceKHR-instance-parameter"
  id="VUID-vkCreateAndroidSurfaceKHR-instance-parameter"></a>
  VUID-vkCreateAndroidSurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateAndroidSurfaceKHR-pCreateInfo-parameter"
  id="VUID-vkCreateAndroidSurfaceKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateAndroidSurfaceKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkAndroidSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateAndroidSurfaceKHR-pAllocator-parameter"
  id="VUID-vkCreateAndroidSurfaceKHR-pAllocator-parameter"></a>
  VUID-vkCreateAndroidSurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateAndroidSurfaceKHR-pSurface-parameter"
  id="VUID-vkCreateAndroidSurfaceKHR-pSurface-parameter"></a>
  VUID-vkCreateAndroidSurfaceKHR-pSurface-parameter  
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

[VK_KHR_android_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_android_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkAndroidSurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateInfoKHR.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateAndroidSurfaceKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
