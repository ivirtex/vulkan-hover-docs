# VkSurfaceProtectedCapabilitiesKHR(3) Manual Page

## Name

VkSurfaceProtectedCapabilitiesKHR - Structure describing capability of a
surface to be protected



## <a href="#_c_specification" class="anchor"></a>C Specification

An application queries if a protected [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)
is displayable on a specific windowing system using
`VkSurfaceProtectedCapabilitiesKHR`, which **can** be passed in `pNext`
parameter of `VkSurfaceCapabilities2KHR`.

The `VkSurfaceProtectedCapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_surface_protected_capabilities
typedef struct VkSurfaceProtectedCapabilitiesKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           supportsProtected;
} VkSurfaceProtectedCapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `supportsProtected` specifies whether a protected swapchain created
  from
  [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface`
  for a particular windowing system **can** be displayed on screen or
  not. If `supportsProtected` is `VK_TRUE`, then creation of swapchains
  with the `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` flag set **must** be
  supported for `surface`.

## <a href="#_description" class="anchor"></a>Description

If the [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
extension is enabled, the value returned in `supportsProtected` will be
identical for every valid surface created on this physical device, and
so in the
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html)
call,
[VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface`
**can** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html). In that case, the
contents of
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html)::`surfaceCapabilities`
as well as any other struct chained to it will be undefined.

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceProtectedCapabilitiesKHR-sType-sType"
  id="VUID-VkSurfaceProtectedCapabilitiesKHR-sType-sType"></a>
  VUID-VkSurfaceProtectedCapabilitiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface_protected_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface_protected_capabilities.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceProtectedCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
