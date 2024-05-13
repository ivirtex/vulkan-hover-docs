# VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR - Structure describing
whether subgroup rotation is enabled



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_shader_subgroup_rotate
typedef struct VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderSubgroupRotate;
    VkBool32           shaderSubgroupRotateClustered;
} VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-shaderSubgroupRotate"></span>
  `shaderSubgroupRotate` specifies whether shader modules **can**
  declare the `GroupNonUniformRotateKHR` capability.

- <span id="features-shaderSubgroupRotateClustered"></span>
  `shaderSubgroupRotateClustered` specifies whether shader modules
  **can** use the `ClusterSize` operand to `OpGroupNonUniformRotateKHR`.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SUBGROUP_ROTATE_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_subgroup_rotate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_rotate.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
