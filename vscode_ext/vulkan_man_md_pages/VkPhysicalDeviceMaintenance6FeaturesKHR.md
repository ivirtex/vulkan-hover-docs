# VkPhysicalDeviceMaintenance6FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance6FeaturesKHR - Structure describing whether
the implementation supports maintenance6 functionality



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance6FeaturesKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6
typedef struct VkPhysicalDeviceMaintenance6FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance6;
} VkPhysicalDeviceMaintenance6FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-maintenance6"></span> `maintenance6` indicates that
  the implementation supports the following:

  - [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **can** be used when binding
    an index buffer

  - [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindMemoryStatusKHR.html) **can** be
    included in the `pNext` chain of the
    [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) and
    [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) structures,
    enabling applications to retrieve [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html) values
    for individual memory binding operations.

  - [VkPhysicalDeviceMaintenance6PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6PropertiesKHR.html)::`blockTexelViewCompatibleMultipleLayers`
    property to indicate that the implementation supports creating image
    views with `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT` where
    the `layerCount` member of `subresourceRange` is greater than `1`.

  - [VkPhysicalDeviceMaintenance6PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6PropertiesKHR.html)::`maxCombinedImageSamplerDescriptorCount`
    property which indicates the maximum descriptor size required for
    any <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
    target="_blank" rel="noopener">format that requires a sampler
    Y′C<sub>B</sub>C<sub>R</sub> conversion</a> supported by the
    implementation.

  - A
    [VkPhysicalDeviceMaintenance6PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6PropertiesKHR.html)::`fragmentShadingRateClampCombinerInputs`
    property which indicates whether the implementation clamps the
    inputs to fragment shading rate combiner operations.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMaintenance6FeaturesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMaintenance6FeaturesKHR` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance6FeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance6FeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance6FeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_6_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance6FeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
