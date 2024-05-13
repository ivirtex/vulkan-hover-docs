# vkInitializePerformanceApiINTEL(3) Manual Page

## Name

vkInitializePerformanceApiINTEL - Initialize a device for performance
queries



## <a href="#_c_specification" class="anchor"></a>C Specification

Prior to creating a performance query pool, initialize the device for
performance queries with the call:

``` c
// Provided by VK_INTEL_performance_query
VkResult vkInitializePerformanceApiINTEL(
    VkDevice                                    device,
    const VkInitializePerformanceApiInfoINTEL*  pInitializeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device used for the queries.

- `pInitializeInfo` is a pointer to a
  [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInitializePerformanceApiInfoINTEL.html)
  structure specifying initialization parameters.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkInitializePerformanceApiINTEL-device-parameter"
  id="VUID-vkInitializePerformanceApiINTEL-device-parameter"></a>
  VUID-vkInitializePerformanceApiINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkInitializePerformanceApiINTEL-pInitializeInfo-parameter"
  id="VUID-vkInitializePerformanceApiINTEL-pInitializeInfo-parameter"></a>
  VUID-vkInitializePerformanceApiINTEL-pInitializeInfo-parameter  
  `pInitializeInfo` **must** be a valid pointer to a valid
  [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInitializePerformanceApiInfoINTEL.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInitializePerformanceApiInfoINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkInitializePerformanceApiINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
