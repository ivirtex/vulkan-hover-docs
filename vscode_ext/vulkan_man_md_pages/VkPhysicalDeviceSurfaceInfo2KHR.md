# VkPhysicalDeviceSurfaceInfo2KHR(3) Manual Page

## Name

VkPhysicalDeviceSurfaceInfo2KHR - Structure specifying a surface and
related swapchain creation parameters



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSurfaceInfo2KHR` structure is defined as:

``` c
// Provided by VK_KHR_get_surface_capabilities2
typedef struct VkPhysicalDeviceSurfaceInfo2KHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSurfaceKHR       surface;
} VkPhysicalDeviceSurfaceInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `surface` is the surface that will be associated with the swapchain.

## <a href="#_description" class="anchor"></a>Description

The members of `VkPhysicalDeviceSurfaceInfo2KHR` correspond to the
arguments to
[vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html),
with `sType` and `pNext` added for extensibility.

Additional capabilities of a surface **may** be available to swapchains
created with different full-screen exclusive settings - particularly if
exclusive full-screen access is application controlled. These additional
capabilities **can** be queried by adding a
[VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)
structure to the `pNext` chain of this structure when used to query
surface properties. Additionally, for Win32 surfaces with application
controlled exclusive full-screen access, chaining a
[VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html)
structure **may** also report additional surface capabilities. These
additional capabilities only apply to swapchains created with the same
parameters included in the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html).

Valid Usage

- <a href="#VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-02672"
  id="VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-02672"></a>
  VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-02672  
  If the `pNext` chain includes a
  [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html)
  structure with its `fullScreenExclusive` member set to
  `VK_FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT`, and `surface`
  was created using
  [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html), a
  [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html)
  structure **must** be included in the `pNext` chain

- <a href="#VUID-VkPhysicalDeviceSurfaceInfo2KHR-surface-07919"
  id="VUID-VkPhysicalDeviceSurfaceInfo2KHR-surface-07919"></a>
  VUID-VkPhysicalDeviceSurfaceInfo2KHR-surface-07919  
  If surface is not VK_NULL_HANDLE, and the
  [`VK_GOOGLE_surfaceless_query`](VK_GOOGLE_surfaceless_query.html)
  extension is not enabled, `surface` **must** be a valid
  [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-sType"
  id="VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-sType"></a>
  VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SURFACE_INFO_2_KHR`

- <a href="#VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-pNext"
  id="VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-pNext"></a>
  VUID-VkPhysicalDeviceSurfaceInfo2KHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkSurfaceFullScreenExclusiveInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveInfoEXT.html),
  [VkSurfaceFullScreenExclusiveWin32InfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFullScreenExclusiveWin32InfoEXT.html),
  or [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a href="#VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-unique"
  id="VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-unique"></a>
  VUID-VkPhysicalDeviceSurfaceInfo2KHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html),
[vkGetDeviceGroupSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupSurfacePresentModes2EXT.html),
[vkGetPhysicalDeviceSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2KHR.html),
[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html),
[vkGetPhysicalDeviceSurfacePresentModes2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfacePresentModes2EXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSurfaceInfo2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
