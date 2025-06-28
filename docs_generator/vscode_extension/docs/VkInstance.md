# VkInstance(3) Manual Page

## Name

VkInstance - Opaque handle to an instance object



## [](#_c_specification)C Specification

There is no global state in Vulkan and all per-application state is stored in a `VkInstance` object. Creating a `VkInstance` object initializes the Vulkan library and allows the application to pass information about itself to the implementation.

Instances are represented by `VkInstance` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_HANDLE(VkInstance)
```

## [](#_see_also)See Also

[PFN\_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkGetInstanceProcAddrLUNARG.html), [VK\_DEFINE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateAndroidSurfaceKHR.html), [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugReportCallbackEXT.html), [vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugUtilsMessengerEXT.html), [vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDirectFBSurfaceEXT.html), [vkCreateDisplayPlaneSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDisplayPlaneSurfaceKHR.html), [vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateHeadlessSurfaceEXT.html), [vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIOSSurfaceMVK.html), [vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImagePipeSurfaceFUCHSIA.html), [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html), [vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMacOSSurfaceMVK.html), [vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMetalSurfaceEXT.html), [vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateScreenSurfaceQNX.html), [vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateStreamDescriptorSurfaceGGP.html), [vkCreateSurfaceOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSurfaceOHOS.html), [vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateViSurfaceNN.html), [vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWaylandSurfaceKHR.html), [vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateWin32SurfaceKHR.html), [vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXcbSurfaceKHR.html), [vkCreateXlibSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateXlibSurfaceKHR.html), [vkDebugReportMessageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDebugReportMessageEXT.html), [vkDestroyDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDebugReportCallbackEXT.html), [vkDestroyDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDebugUtilsMessengerEXT.html), [vkDestroyInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyInstance.html), [vkDestroySurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySurfaceKHR.html), [vkEnumeratePhysicalDeviceGroups](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceGroups.html), [vkEnumeratePhysicalDeviceGroupsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceGroupsKHR.html), [vkEnumeratePhysicalDevices](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDevices.html), [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html), [vkSubmitDebugUtilsMessageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSubmitDebugUtilsMessageEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkInstance)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0