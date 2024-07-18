# VkPhysicalDeviceSubgroupSizeControlProperties(3) Manual Page

## Name

VkPhysicalDeviceSubgroupSizeControlProperties - Structure describing the
control subgroup size properties of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSubgroupSizeControlProperties` structure is defined
as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceSubgroupSizeControlProperties {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              minSubgroupSize;
    uint32_t              maxSubgroupSize;
    uint32_t              maxComputeWorkgroupSubgroups;
    VkShaderStageFlags    requiredSubgroupSizeStages;
} VkPhysicalDeviceSubgroupSizeControlProperties;
```

or the equivalent

``` c
// Provided by VK_EXT_subgroup_size_control
typedef VkPhysicalDeviceSubgroupSizeControlProperties VkPhysicalDeviceSubgroupSizeControlPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-minSubgroupSize"></span> `minSubgroupSize`
  is the minimum subgroup size supported by this device.
  `minSubgroupSize` is at least one if any of the physical device’s
  queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`.
  `minSubgroupSize` is a power-of-two. `minSubgroupSize` is less than or
  equal to `maxSubgroupSize`. `minSubgroupSize` is less than or equal to
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupSize"
  target="_blank" rel="noopener"><code>subgroupSize</code></a>.

- <span id="extension-limits-maxSubgroupSize"></span> `maxSubgroupSize`
  is the maximum subgroup size supported by this device.
  `maxSubgroupSize` is at least one if any of the physical device’s
  queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`.
  `maxSubgroupSize` is a power-of-two. `maxSubgroupSize` is greater than
  or equal to `minSubgroupSize`. `maxSubgroupSize` is greater than or
  equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupSize"
  target="_blank" rel="noopener"><code>subgroupSize</code></a>.

- <span id="extension-limits-maxComputeWorkgroupSubgroups"></span>
  `maxComputeWorkgroupSubgroups` is the maximum number of subgroups
  supported by the implementation within a workgroup.

- <span id="extension-limits-requiredSubgroupSizeStages"></span>
  `requiredSubgroupSizeStages` is a bitfield of what shader stages
  support having a required subgroup size specified.

If the `VkPhysicalDeviceSubgroupSizeControlProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

If
[VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupProperties.html)::`supportedOperations`
includes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroup-quad"
target="_blank"
rel="noopener"><code>VK_SUBGROUP_FEATURE_QUAD_BIT</code></a>,
`minSubgroupSize` **must** be greater than or equal to 4.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceSubgroupSizeControlProperties-sType-sType"
  id="VUID-VkPhysicalDeviceSubgroupSizeControlProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceSubgroupSizeControlProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subgroup_size_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subgroup_size_control.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSubgroupSizeControlProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
