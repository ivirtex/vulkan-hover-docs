# VkPhysicalDeviceVulkanMemoryModelFeatures(3) Manual Page

## Name

VkPhysicalDeviceVulkanMemoryModelFeatures - Structure describing
features supported by the memory model



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVulkanMemoryModelFeatures` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceVulkanMemoryModelFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           vulkanMemoryModel;
    VkBool32           vulkanMemoryModelDeviceScope;
    VkBool32           vulkanMemoryModelAvailabilityVisibilityChains;
} VkPhysicalDeviceVulkanMemoryModelFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_vulkan_memory_model
typedef VkPhysicalDeviceVulkanMemoryModelFeatures VkPhysicalDeviceVulkanMemoryModelFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-vulkanMemoryModel"></span>
  `vulkanMemoryModel` indicates whether shader modules **can** declare
  the `VulkanMemoryModel` capability.

- <span id="extension-features-vulkanMemoryModelDeviceScope"></span>
  `vulkanMemoryModelDeviceScope` indicates whether the Vulkan Memory
  Model can use `Device` scope synchronization. This also indicates
  whether shader modules **can** declare the
  `VulkanMemoryModelDeviceScope` capability.

- <span id="extension-features-vulkanMemoryModelAvailabilityVisibilityChains"></span>
  `vulkanMemoryModelAvailabilityVisibilityChains` indicates whether the
  Vulkan Memory Model can use <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-model-availability-visibility"
  target="_blank" rel="noopener">availability and visibility chains</a>
  with more than one element.

If the `VkPhysicalDeviceVulkanMemoryModelFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVulkanMemoryModelFeaturesKHR` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVulkanMemoryModelFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceVulkanMemoryModelFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceVulkanMemoryModelFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVulkanMemoryModelFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
