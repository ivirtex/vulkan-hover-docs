# VkSurfaceProtectedCapabilitiesKHR(3) Manual Page

## Name

VkSurfaceProtectedCapabilitiesKHR - Structure describing capability of a surface to be protected



## [](#_c_specification)C Specification

An application queries if a protected [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) is displayable on a specific windowing system using `VkSurfaceProtectedCapabilitiesKHR`, which **can** be passed in `pNext` parameter of `VkSurfaceCapabilities2KHR`.

The `VkSurfaceProtectedCapabilitiesKHR` structure is defined as:

```c++
// Provided by VK_KHR_surface_protected_capabilities
typedef struct VkSurfaceProtectedCapabilitiesKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           supportsProtected;
} VkSurfaceProtectedCapabilitiesKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `supportsProtected` specifies whether a protected swapchain created from [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface` for a particular windowing system **can** be displayed on screen or not. If `supportsProtected` is `VK_TRUE`, then creation of swapchains with the `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` flag set **must** be supported for `surface`.

## [](#_description)Description

If the `VK_GOOGLE_surfaceless_query` extension is enabled, the value returned in `supportsProtected` will be identical for every valid surface created on this physical device, and so in the [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html) call, [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html)::`surface` **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). In that case, the contents of [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html)::`surfaceCapabilities` as well as any other structure chained to it will be undefined.

Valid Usage (Implicit)

- [](#VUID-VkSurfaceProtectedCapabilitiesKHR-sType-sType)VUID-VkSurfaceProtectedCapabilitiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_PROTECTED_CAPABILITIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_surface\_protected\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface_protected_capabilities.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceProtectedCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0