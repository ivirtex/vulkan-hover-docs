# vkDestroySurfaceKHR(3) Manual Page

## Name

vkDestroySurfaceKHR - Destroy a VkSurfaceKHR object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a `VkSurfaceKHR` object, call:

``` c
// Provided by VK_KHR_surface
void vkDestroySurfaceKHR(
    VkInstance                                  instance,
    VkSurfaceKHR                                surface,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance used to create the surface.

- `surface` is the surface to destroy.

- `pAllocator` is the allocator used for host memory allocated for the
  surface object when there is no more specific allocator available (see
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

## <a href="#_description" class="anchor"></a>Description

Destroying a `VkSurfaceKHR` merely severs the connection between Vulkan
and the native surface, and does not imply destroying the native
surface, closing a window, or similar behavior.

Valid Usage

- <a href="#VUID-vkDestroySurfaceKHR-surface-01266"
  id="VUID-vkDestroySurfaceKHR-surface-01266"></a>
  VUID-vkDestroySurfaceKHR-surface-01266  
  All `VkSwapchainKHR` objects created for `surface` **must** have been
  destroyed prior to destroying `surface`

- <a href="#VUID-vkDestroySurfaceKHR-surface-01267"
  id="VUID-vkDestroySurfaceKHR-surface-01267"></a>
  VUID-vkDestroySurfaceKHR-surface-01267  
  If `VkAllocationCallbacks` were provided when `surface` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroySurfaceKHR-surface-01268"
  id="VUID-vkDestroySurfaceKHR-surface-01268"></a>
  VUID-vkDestroySurfaceKHR-surface-01268  
  If no `VkAllocationCallbacks` were provided when `surface` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroySurfaceKHR-instance-parameter"
  id="VUID-vkDestroySurfaceKHR-instance-parameter"></a>
  VUID-vkDestroySurfaceKHR-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkDestroySurfaceKHR-surface-parameter"
  id="VUID-vkDestroySurfaceKHR-surface-parameter"></a>
  VUID-vkDestroySurfaceKHR-surface-parameter  
  If `surface` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `surface`
  **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkDestroySurfaceKHR-pAllocator-parameter"
  id="VUID-vkDestroySurfaceKHR-pAllocator-parameter"></a>
  VUID-vkDestroySurfaceKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroySurfaceKHR-surface-parent"
  id="VUID-vkDestroySurfaceKHR-surface-parent"></a>
  VUID-vkDestroySurfaceKHR-surface-parent  
  If `surface` is a valid handle, it **must** have been created,
  allocated, or retrieved from `instance`

Host Synchronization

- Host access to `surface` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroySurfaceKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
