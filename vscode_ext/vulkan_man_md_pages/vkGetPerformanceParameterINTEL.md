# vkGetPerformanceParameterINTEL(3) Manual Page

## Name

vkGetPerformanceParameterINTEL - Query performance capabilities of the
device



## <a href="#_c_specification" class="anchor"></a>C Specification

Some performance query features of a device can be discovered with the
call:

``` c
// Provided by VK_INTEL_performance_query
VkResult vkGetPerformanceParameterINTEL(
    VkDevice                                    device,
    VkPerformanceParameterTypeINTEL             parameter,
    VkPerformanceValueINTEL*                    pValue);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device to query.

- `parameter` is the parameter to query.

- `pValue` is a pointer to a
  [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html) structure in
  which the type and value of the parameter are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetPerformanceParameterINTEL-device-parameter"
  id="VUID-vkGetPerformanceParameterINTEL-device-parameter"></a>
  VUID-vkGetPerformanceParameterINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetPerformanceParameterINTEL-parameter-parameter"
  id="VUID-vkGetPerformanceParameterINTEL-parameter-parameter"></a>
  VUID-vkGetPerformanceParameterINTEL-parameter-parameter  
  `parameter` **must** be a valid
  [VkPerformanceParameterTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceParameterTypeINTEL.html)
  value

- <a href="#VUID-vkGetPerformanceParameterINTEL-pValue-parameter"
  id="VUID-vkGetPerformanceParameterINTEL-pValue-parameter"></a>
  VUID-vkGetPerformanceParameterINTEL-pValue-parameter  
  `pValue` **must** be a valid pointer to a
  [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPerformanceParameterTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceParameterTypeINTEL.html),
[VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPerformanceParameterINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
