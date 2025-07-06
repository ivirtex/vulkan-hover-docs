# VkPhysicalDeviceSurfaceInfo2KHR(3) Manual Page

## Name

VkPhysicalDeviceSurfaceInfo2KHR - Structure specifying a surface and related swapchain creation parameters



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSurfaceInfo2KHR` structure is defined as:

```c++
// Provided by VK_KHR_get_surface_capabilities2
typedef struct VkPhysicalDeviceSurfaceInfo2KHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSurfaceKHR       surface;
} VkPhysicalDeviceSurfaceInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `surface` is the surface that will be associated with the swapchain.

## [](#_description)Description

The members of `VkPhysicalDeviceSurfaceInfo2KHR` correspond to the arguments to [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html), with `sType` and `pNext` added for extensibility.

Additional capabilities of a surface **may** be available to swapchains created with different full-screen exclusive settings - particularly if exclusive full-screen access is application controlled. These additional capabilities **can** be queried by adding a [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html) structure to the `pNext` chain of this structure when used to query surface properties. Additionally, for Win32 surfaces with application controlled exclusive full-screen access, chaining a [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html) structure **may** also report additional surface capabilities. These additional capabilities only apply to swapchains created with the same parameters included in the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html).

Valid Usage

- [](#VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-02672)VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-02672  
  If the `pNext` chain includes a [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html) structure with its `fullScreenExclusive` member set to `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`, and `surface` was created using [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWin32SurfaceKHR.html), a [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html) structure **must** be included in the `pNext` chain
- [](#VUID-VkPhysicalDeviceSurfaceInfo2KHR-surface-07919)VUID-VkPhysicalDeviceSurfaceInfo2KHR-surface-07919  
  If surface is not VK\_NULL\_HANDLE, and the `VK_GOOGLE_surfaceless_query` extension is not enabled, `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) handle

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-sType)VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR`
- [](#VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-pNext)VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html), [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html), or [VkSurfacePresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfacePresentModeKHR.html)
- [](#VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-unique)VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html), [vkGetDeviceGroupSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html), [vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html), [vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html), [vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSurfaceInfo2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0