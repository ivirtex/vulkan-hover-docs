# VkPerformanceConfigurationINTEL(3) Manual Page

## Name

VkPerformanceConfigurationINTEL - Device configuration for performance
queries



## <a href="#_c_specification" class="anchor"></a>C Specification

Before submitting command buffers containing performance queries
commands to a device queue, the application must acquire and set a
performance query configuration. The configuration can be released once
all command buffers containing performance query commands are not in a
pending state.

``` c
// Provided by VK_INTEL_performance_query
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPerformanceConfigurationINTEL)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_INTEL_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_INTEL_performance_query.html),
[vkAcquirePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquirePerformanceConfigurationINTEL.html),
[vkQueueSetPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSetPerformanceConfigurationINTEL.html),
[vkReleasePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleasePerformanceConfigurationINTEL.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceConfigurationINTEL"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
