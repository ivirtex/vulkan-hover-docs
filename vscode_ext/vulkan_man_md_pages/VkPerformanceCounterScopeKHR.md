# VkPerformanceCounterScopeKHR(3) Manual Page

## Name

VkPerformanceCounterScopeKHR - Supported counter scope types



## <a href="#_c_specification" class="anchor"></a>C Specification

Performance counters have an associated scope. This scope describes the
granularity of a performance counter.

The performance counter scope types which **may** be returned in
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)::`scope` are:

``` c
// Provided by VK_KHR_performance_query
typedef enum VkPerformanceCounterScopeKHR {
    VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR = 0,
    VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR = 1,
    VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR = 2,
    VK_QUERY_SCOPE_COMMAND_BUFFER_KHR = VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR,
    VK_QUERY_SCOPE_RENDER_PASS_KHR = VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR,
    VK_QUERY_SCOPE_COMMAND_KHR = VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR,
} VkPerformanceCounterScopeKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_BUFFER_KHR` - the performance
  counter scope is a single complete command buffer.

- `VK_PERFORMANCE_COUNTER_SCOPE_RENDER_PASS_KHR` - the performance
  counter scope is zero or more complete render passes. The performance
  query containing the performance counter **must** begin and end
  outside a render pass instance.

- `VK_PERFORMANCE_COUNTER_SCOPE_COMMAND_KHR` - the performance counter
  scope is zero or more commands.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_performance_query](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_performance_query.html),
[VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPerformanceCounterScopeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
