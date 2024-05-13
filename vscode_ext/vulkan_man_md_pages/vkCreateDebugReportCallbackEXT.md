# vkCreateDebugReportCallbackEXT(3) Manual Page

## Name

vkCreateDebugReportCallbackEXT - Create a debug report callback object



## <a href="#_c_specification" class="anchor"></a>C Specification

Debug report callbacks give more detailed feedback on the application’s
use of Vulkan when events of interest occur.

To register a debug report callback, an application uses
[vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugReportCallbackEXT.html).

``` c
// Provided by VK_EXT_debug_report
VkResult vkCreateDebugReportCallbackEXT(
    VkInstance                                  instance,
    const VkDebugReportCallbackCreateInfoEXT*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDebugReportCallbackEXT*                   pCallback);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance the callback will be logged on.

- `pCreateInfo` is a pointer to a
  [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)
  structure defining the conditions under which this callback will be
  called.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pCallback` is a pointer to a
  [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html) handle in
  which the created object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDebugReportCallbackEXT-instance-parameter"
  id="VUID-vkCreateDebugReportCallbackEXT-instance-parameter"></a>
  VUID-vkCreateDebugReportCallbackEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateDebugReportCallbackEXT-pCreateInfo-parameter"
  id="VUID-vkCreateDebugReportCallbackEXT-pCreateInfo-parameter"></a>
  VUID-vkCreateDebugReportCallbackEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html)
  structure

- <a href="#VUID-vkCreateDebugReportCallbackEXT-pAllocator-parameter"
  id="VUID-vkCreateDebugReportCallbackEXT-pAllocator-parameter"></a>
  VUID-vkCreateDebugReportCallbackEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDebugReportCallbackEXT-pCallback-parameter"
  id="VUID-vkCreateDebugReportCallbackEXT-pCallback-parameter"></a>
  VUID-vkCreateDebugReportCallbackEXT-pCallback-parameter  
  `pCallback` **must** be a valid pointer to a
  [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackCreateInfoEXT.html),
[VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDebugReportCallbackEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
