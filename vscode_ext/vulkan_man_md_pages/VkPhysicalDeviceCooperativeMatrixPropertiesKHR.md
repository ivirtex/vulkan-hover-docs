# VkPhysicalDeviceCooperativeMatrixPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceCooperativeMatrixPropertiesKHR - Structure describing
cooperative matrix properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCooperativeMatrixPropertiesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_cooperative_matrix
typedef struct VkPhysicalDeviceCooperativeMatrixPropertiesKHR {
    VkStructureType       sType;
    void*                 pNext;
    VkShaderStageFlags    cooperativeMatrixSupportedStages;
} VkPhysicalDeviceCooperativeMatrixPropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-cooperativeMatrixSupportedStages"></span>
  `cooperativeMatrixSupportedStages` is a bitfield of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) describing the
  shader stages that cooperative matrix instructions are supported in.
  `cooperativeMatrixSupportedStages` will have the
  `VK_SHADER_STAGE_COMPUTE_BIT` bit set if any of the physical device’s
  queues support `VK_QUEUE_COMPUTE_BIT`.

## <a href="#_description" class="anchor"></a>Description

`cooperativeMatrixSupportedStages` **must** not have any bits other than
`VK_SHADER_STAGE_COMPUTE_BIT` set.

If the `VkPhysicalDeviceCooperativeMatrixPropertiesKHR` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceCooperativeMatrixPropertiesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceCooperativeMatrixPropertiesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceCooperativeMatrixPropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_cooperative_matrix](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_cooperative_matrix.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCooperativeMatrixPropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
