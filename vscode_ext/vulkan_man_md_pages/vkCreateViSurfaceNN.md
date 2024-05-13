# vkCreateViSurfaceNN(3) Manual Page

## Name

vkCreateViSurfaceNN - Create a slink:VkSurfaceKHR object for a VI layer



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a `VkSurfaceKHR` object for an `nn`::`vi`::`Layer`, query the
layer’s native handle using `nn`::`vi`::`GetNativeWindow`, and then
call:

``` c
// Provided by VK_NN_vi_surface
VkResult vkCreateViSurfaceNN(
    VkInstance                                  instance,
    const VkViSurfaceCreateInfoNN*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance with which to associate the surface.

- `pCreateInfo` is a pointer to a `VkViSurfaceCreateInfoNN` structure
  containing parameters affecting the creation of the surface object.

- `pAllocator` is the allocator used for host memory allocated for the
  surface object when there is no more specific allocator available (see
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

- `pSurface` is a pointer to a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle
  in which the created surface object is returned.

## <a href="#_description" class="anchor"></a>Description

During the lifetime of a surface created using a particular
`nn`::`vi`::`NativeWindowHandle`, applications **must** not attempt to
create another surface for the same `nn`::`vi`::`Layer` or attempt to
connect to the same `nn`::`vi`::`Layer` through other platform
mechanisms.

If the native window is created with a specified size, `currentExtent`
will reflect that size. In this case, applications should use the same
size for the swapchain’s `imageExtent`. Otherwise, the `currentExtent`
will have the special value (0xFFFFFFFF, 0xFFFFFFFF), indicating that
applications are expected to choose an appropriate size for the
swapchain’s `imageExtent` (e.g., by matching the result of a call to
`nn`::`vi`::`GetDisplayResolution`).

Valid Usage (Implicit)

- <a href="#VUID-vkCreateViSurfaceNN-instance-parameter"
  id="VUID-vkCreateViSurfaceNN-instance-parameter"></a>
  VUID-vkCreateViSurfaceNN-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateViSurfaceNN-pCreateInfo-parameter"
  id="VUID-vkCreateViSurfaceNN-pCreateInfo-parameter"></a>
  VUID-vkCreateViSurfaceNN-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateInfoNN.html) structure

- <a href="#VUID-vkCreateViSurfaceNN-pAllocator-parameter"
  id="VUID-vkCreateViSurfaceNN-pAllocator-parameter"></a>
  VUID-vkCreateViSurfaceNN-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateViSurfaceNN-pSurface-parameter"
  id="VUID-vkCreateViSurfaceNN-pSurface-parameter"></a>
  VUID-vkCreateViSurfaceNN-pSurface-parameter  
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

[VK_NN_vi_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NN_vi_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html),
[VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateInfoNN.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateViSurfaceNN"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
