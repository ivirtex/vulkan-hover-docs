# VkInstance(3) Manual Page

## Name

VkInstance - Opaque handle to an instance object



## <a href="#_c_specification" class="anchor"></a>C Specification

There is no global state in Vulkan and all per-application state is
stored in a `VkInstance` object. Creating a `VkInstance` object
initializes the Vulkan library and allows the application to pass
information about itself to the implementation.

Instances are represented by `VkInstance` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_HANDLE(VkInstance)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAndroidSurfaceKHR.html),
[vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugReportCallbackEXT.html),
[vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html),
[vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDirectFBSurfaceEXT.html),
[vkCreateDisplayPlaneSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDisplayPlaneSurfaceKHR.html),
[vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateHeadlessSurfaceEXT.html),
[vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIOSSurfaceMVK.html),
[vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImagePipeSurfaceFUCHSIA.html),
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html),
[vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMacOSSurfaceMVK.html),
[vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMetalSurfaceEXT.html),
[vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateScreenSurfaceQNX.html),
[vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateStreamDescriptorSurfaceGGP.html),
[vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateViSurfaceNN.html),
[vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWaylandSurfaceKHR.html),
[vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html),
[vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXcbSurfaceKHR.html),
[vkCreateXlibSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXlibSurfaceKHR.html),
[vkDebugReportMessageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDebugReportMessageEXT.html),
[vkDestroyDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugReportCallbackEXT.html),
[vkDestroyDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugUtilsMessengerEXT.html),
[vkDestroyInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyInstance.html),
[vkDestroySurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySurfaceKHR.html),
[vkEnumeratePhysicalDeviceGroups](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroups.html),
[vkEnumeratePhysicalDeviceGroupsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceGroupsKHR.html),
[vkEnumeratePhysicalDevices](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDevices.html),
[vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetInstanceProcAddr.html),
[vkSubmitDebugUtilsMessageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSubmitDebugUtilsMessageEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkInstance"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
