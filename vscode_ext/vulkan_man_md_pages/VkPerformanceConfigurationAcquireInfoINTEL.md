# VkPerformanceConfigurationAcquireInfoINTEL(3) Manual Page

## Name

VkPerformanceConfigurationAcquireInfoINTEL - Acquire a configuration to
capture performance data



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPerformanceConfigurationAcquireInfoINTEL` structure is defined
as:

``` c
// Provided by VK_INTEL_performance_query
typedef struct VkPerformanceConfigurationAcquireInfoINTEL {
    VkStructureType                        sType;
    const void*                            pNext;
    VkPerformanceConfigurationTypeINTEL    type;
} VkPerformanceConfigurationAcquireInfoINTEL;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` is one of the
  [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationTypeINTEL.html)
  type of performance configuration that will be acquired.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPerformanceConfigurationAcquireInfoINTEL-sType-sType"
  id="VUID-VkPerformanceConfigurationAcquireInfoINTEL-sType-sType"></a>
  VUID-VkPerformanceConfigurationAcquireInfoINTEL-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL`

- <a href="#VUID-VkPerformanceConfigurationAcquireInfoINTEL-pNext-pNext"
  id="VUID-VkPerformanceConfigurationAcquireInfoINTEL-pNext-pNext"></a>
  VUID-VkPerformanceConfigurationAcquireInfoINTEL-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkPerformanceConfigurationAcquireInfoINTEL-type-parameter"
  id="VUID-VkPerformanceConfigurationAcquireInfoINTEL-type-parameter"></a>
  VUID-VkPerformanceConfigurationAcquireInfoINTEL-type-parameter  
  `type` **must** be a valid
  [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationTypeINTEL.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationTypeINTEL.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkAcquirePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquirePerformanceConfigurationINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceConfigurationAcquireInfoINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
