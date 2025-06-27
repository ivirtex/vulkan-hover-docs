# VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT - Structure describing
whether implementation supports swapchain maintenance1 functionality



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           swapchainMaintenance1;
} VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-swapchainMaintenance1"></span>
  `swapchainMaintenance1` indicates that the implementation supports the
  following:

  - [VkSwapchainPresentFenceInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentFenceInfoEXT.html),
    specifying a fence that is signaled when the resources associated
    with a present operation **can** be safely destroyed.

  - [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)
    and
    [VkSwapchainPresentModeInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModeInfoEXT.html),
    allowing the swapchain to switch present modes without a need for
    recreation.

  - [VkSwapchainPresentScalingCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentScalingCreateInfoEXT.html),
    specifying the scaling behavior of the swapchain in presence of
    window resizing.

  - The `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT` flag,
    allowing the implementation to defer the allocation of swapchain
    image memory until first acquisition.

  - [vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseSwapchainImagesEXT.html),
    allowing acquired swapchain images to be released without presenting
    them.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SWAPCHAIN_MAINTENANCE_1_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
