# vkCreateWin32SurfaceKHR(3) Manual Page

## Name

vkCreateWin32SurfaceKHR - Create a VkSurfaceKHR object for an Win32
native window



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a `VkSurfaceKHR` object for a Win32 window, call:

``` c
// Provided by VK_KHR_win32_surface
VkResult vkCreateWin32SurfaceKHR(
    VkInstance                                  instance,
    const VkWin32SurfaceCreateInfoKHR*          pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance to associate the surface with.

- `pCreateInfo` is a pointer to a `VkWin32SurfaceCreateInfoKHR`
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

Valid Usage (Implicit)

- <a href="#VUID-vkCreateWin32SurfaceKHR-instance-parameter"
  id="VUID-vkCreateWin32SurfaceKHR-instance-parameter"></a>
  VUID-vkCreateWin32SurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateWin32SurfaceKHR-pCreateInfo-parameter"
  id="VUID-vkCreateWin32SurfaceKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateWin32SurfaceKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkWin32SurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateWin32SurfaceKHR-pAllocator-parameter"
  id="VUID-vkCreateWin32SurfaceKHR-pAllocator-parameter"></a>
  VUID-vkCreateWin32SurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateWin32SurfaceKHR-pSurface-parameter"
  id="VUID-vkCreateWin32SurfaceKHR-pSurface-parameter"></a>
  VUID-vkCreateWin32SurfaceKHR-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

Some Vulkan functions **may** call the `SendMessage` system API when
interacting with a `VkSurfaceKHR` through a `VkSwapchainKHR`. In a
multithreaded environment, calling `SendMessage` from a thread that is
not the thread associated with `pCreateInfo`::`hwnd` will block until
the application has processed the window message. Thus, applications
**should** either call these Vulkan functions on the message pump
thread, or make sure their message pump is actively running. Failing to
do so **may** result in deadlocks.

The functions subject to this requirement are:

- [vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html)

- [vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html)

- [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html) and
  [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html)

- [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueuePresentKHR.html)

- [vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseSwapchainImagesEXT.html)

- [vkAcquireFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireFullScreenExclusiveModeEXT.html)

- [vkReleaseFullScreenExclusiveModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseFullScreenExclusiveModeEXT.html)

- [vkSetHdrMetadataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetHdrMetadataEXT.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_win32_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_win32_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html),
[VkWin32SurfaceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateWin32SurfaceKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
