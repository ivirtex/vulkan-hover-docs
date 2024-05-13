# vkUninitializePerformanceApiINTEL(3) Manual Page

## Name

vkUninitializePerformanceApiINTEL - Uninitialize a device for
performance queries



## <a href="#_c_specification" class="anchor"></a>C Specification

Once performance query operations have completed, uninitialize the
device for performance queries with the call:

``` c
// Provided by VK_INTEL_performance_query
void vkUninitializePerformanceApiINTEL(
    VkDevice                                    device);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device used for the queries.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkUninitializePerformanceApiINTEL-device-parameter"
  id="VUID-vkUninitializePerformanceApiINTEL-device-parameter"></a>
  VUID-vkUninitializePerformanceApiINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkUninitializePerformanceApiINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
