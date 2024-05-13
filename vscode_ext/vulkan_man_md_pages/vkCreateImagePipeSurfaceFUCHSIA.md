# vkCreateImagePipeSurfaceFUCHSIA(3) Manual Page

## Name

vkCreateImagePipeSurfaceFUCHSIA - Create a slink:VkSurfaceKHR object for
a Fuchsia ImagePipe



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a `VkSurfaceKHR` object for a Fuchsia ImagePipe, call:

``` c
// Provided by VK_FUCHSIA_imagepipe_surface
VkResult vkCreateImagePipeSurfaceFUCHSIA(
    VkInstance                                  instance,
    const VkImagePipeSurfaceCreateInfoFUCHSIA*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSurfaceKHR*                               pSurface);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance to associate with the surface.

- `pCreateInfo` is a pointer to a
  [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html)
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

- <a href="#VUID-vkCreateImagePipeSurfaceFUCHSIA-instance-parameter"
  id="VUID-vkCreateImagePipeSurfaceFUCHSIA-instance-parameter"></a>
  VUID-vkCreateImagePipeSurfaceFUCHSIA-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateImagePipeSurfaceFUCHSIA-pCreateInfo-parameter"
  id="VUID-vkCreateImagePipeSurfaceFUCHSIA-pCreateInfo-parameter"></a>
  VUID-vkCreateImagePipeSurfaceFUCHSIA-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html)
  structure

- <a href="#VUID-vkCreateImagePipeSurfaceFUCHSIA-pAllocator-parameter"
  id="VUID-vkCreateImagePipeSurfaceFUCHSIA-pAllocator-parameter"></a>
  VUID-vkCreateImagePipeSurfaceFUCHSIA-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateImagePipeSurfaceFUCHSIA-pSurface-parameter"
  id="VUID-vkCreateImagePipeSurfaceFUCHSIA-pSurface-parameter"></a>
  VUID-vkCreateImagePipeSurfaceFUCHSIA-pSurface-parameter  
  `pSurface` **must** be a valid pointer to a
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_imagepipe_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_imagepipe_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateImagePipeSurfaceFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
