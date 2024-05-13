# vkAcquirePerformanceConfigurationINTEL(3) Manual Page

## Name

vkAcquirePerformanceConfigurationINTEL - Acquire the performance query
capability



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire a device performance configuration, call:

``` c
// Provided by VK_INTEL_performance_query
VkResult vkAcquirePerformanceConfigurationINTEL(
    VkDevice                                    device,
    const VkPerformanceConfigurationAcquireInfoINTEL* pAcquireInfo,
    VkPerformanceConfigurationINTEL*            pConfiguration);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that the performance query commands
  will be submitted to.

- `pAcquireInfo` is a pointer to a
  [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html)
  structure, specifying the performance configuration to acquire.

- `pConfiguration` is a pointer to a `VkPerformanceConfigurationINTEL`
  handle in which the resulting configuration object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkAcquirePerformanceConfigurationINTEL-device-parameter"
  id="VUID-vkAcquirePerformanceConfigurationINTEL-device-parameter"></a>
  VUID-vkAcquirePerformanceConfigurationINTEL-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkAcquirePerformanceConfigurationINTEL-pAcquireInfo-parameter"
  id="VUID-vkAcquirePerformanceConfigurationINTEL-pAcquireInfo-parameter"></a>
  VUID-vkAcquirePerformanceConfigurationINTEL-pAcquireInfo-parameter  
  `pAcquireInfo` **must** be a valid pointer to a valid
  [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html)
  structure

- <a
  href="#VUID-vkAcquirePerformanceConfigurationINTEL-pConfiguration-parameter"
  id="VUID-vkAcquirePerformanceConfigurationINTEL-pConfiguration-parameter"></a>
  VUID-vkAcquirePerformanceConfigurationINTEL-pConfiguration-parameter  
  `pConfiguration` **must** be a valid pointer to a
  [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html)
  handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html),
[VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquirePerformanceConfigurationINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
