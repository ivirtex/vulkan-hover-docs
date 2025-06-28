# VkSurfaceKHR(3) Manual Page

## Name

VkSurfaceKHR - Opaque handle to a surface object



## [](#_c_specification)C Specification

Native platform surface or window objects are abstracted by surface objects, which are represented by `VkSurfaceKHR` handles:

```c++
// Provided by VK_KHR_surface
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSurfaceKHR)
```

## [](#_description)Description

The `VK_KHR_surface` extension declares the `VkSurfaceKHR` object, and provides a function for destroying `VkSurfaceKHR` objects. Separate platform-specific extensions each provide a function for creating a `VkSurfaceKHR` object for the respective platform. From the application’s perspective this is an opaque handle, just like the handles of other Vulkan objects.

Note

On certain platforms, the Vulkan loader and ICDs **may** have conventions that treat the handle as a pointer to a structure containing the platform-specific information about the surface. This will be described in the documentation for the loader-ICD interface, and in the `vk_icd.h` header file of the LoaderAndTools source-code repository. This does not affect the loader-layer interface; layers **may** wrap `VkSurfaceKHR` objects.

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html), [VkPhysicalDeviceSurfaceInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSurfaceInfo2KHR.html), [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html), [vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAndroidSurfaceKHR.html), [vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDirectFBSurfaceEXT.html), [vkCreateDisplayPlaneSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDisplayPlaneSurfaceKHR.html), [vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateHeadlessSurfaceEXT.html), [vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIOSSurfaceMVK.html), [vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImagePipeSurfaceFUCHSIA.html), [vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMacOSSurfaceMVK.html), [vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMetalSurfaceEXT.html), [vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateScreenSurfaceQNX.html), [vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateStreamDescriptorSurfaceGGP.html), [vkCreateSurfaceOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSurfaceOHOS.html), [vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateViSurfaceNN.html), [vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWaylandSurfaceKHR.html), [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWin32SurfaceKHR.html), [vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXcbSurfaceKHR.html), [vkCreateXlibSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXlibSurfaceKHR.html), [vkDestroySurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySurfaceKHR.html), [vkGetDeviceGroupSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceGroupSurfacePresentModesKHR.html), [vkGetPhysicalDevicePresentRectanglesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDevicePresentRectanglesKHR.html), [vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html), [vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html), [vkGetPhysicalDeviceSurfaceFormatsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceFormatsKHR.html), [vkGetPhysicalDeviceSurfacePresentModesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfacePresentModesKHR.html), [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0