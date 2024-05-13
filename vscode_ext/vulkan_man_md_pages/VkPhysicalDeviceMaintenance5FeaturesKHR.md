# VkPhysicalDeviceMaintenance5FeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceMaintenance5FeaturesKHR - Structure describing whether
the implementation supports maintenance5 functionality



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceMaintenance5FeaturesKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkPhysicalDeviceMaintenance5FeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           maintenance5;
} VkPhysicalDeviceMaintenance5FeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-maintenance5"></span> `maintenance5` indicates that
  the implementation supports the following:

  - The ability to expose support for the optional format
    `VK_FORMAT_A1B5G5R5_UNORM_PACK16_KHR`.

  - The ability to expose support for the optional format
    `VK_FORMAT_A8_UNORM_KHR`.

  - A property to indicate that multisample coverage operations are
    performed after sample counting in EarlyFragmentTests mode.

  - Creating a `VkBufferView` with a subset of the associated `VkBuffer`
    usage using
    [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html).

  - A new function
    [vkCmdBindIndexBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindIndexBuffer2KHR.html), allowing
    a range of memory to be bound as an index buffer.

  - [vkGetDeviceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceProcAddr.html) will return `NULL`
    for function pointers of core functions for versions higher than the
    version requested by the application.

  - [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers2.html) supports
    using `VK_WHOLE_SIZE` in the `pSizes` parameter.

  - If `PointSize` is not written, a default value of `1.0` is used for
    the size of points.

  - [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) **can** be
    added as a chained structure to pipeline creation via
    [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
    rather than having to create a shader module.

  - A function
    [vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderingAreaGranularityKHR.html)
    to query the optimal render area for a dynamic rendering instance.

  - A property to indicate that depth/stencil texturing operations with
    `VK_COMPONENT_SWIZZLE_ONE` have defined behavior.

  - [vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSubresourceLayoutKHR.html)
    allows an application to perform a
    [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html)
    query without having to create an image.

  - `VK_REMAINING_ARRAY_LAYERS` as the `layerCount` member of
    [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html).

  - A property to indicate whether `PointSize` controls the final
    rasterization of polygons if <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-polygonmode"
    target="_blank" rel="noopener">polygon mode</a> is
    `VK_POLYGON_MODE_POINT`.

  - Two properties to indicate the non-strict line rasterization
    algorithm used.

  - Two new flags words
    [VkPipelineCreateFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits2KHR.html)
    and [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html).

  - Physical-device-level functions **can** now be called with any value
    in the valid range for a type beyond the defined enumerants, such
    that applications can avoid checking individual features,
    extensions, or versions before querying supported properties of a
    particular enumerant.

  - Copies between images of any type are allowed, with 1D images
    treated as 2D images with a height of `1`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceMaintenance5FeaturesKHR` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceMaintenance5FeaturesKHR` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMaintenance5FeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceMaintenance5FeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceMaintenance5FeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_5_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMaintenance5FeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
