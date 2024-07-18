# VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT(3) Manual Page

## Name

VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT - Structure
describing whether querying shader group handles from a pipeline library
is supported by the implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT` structure
is defined as:

``` c
// Provided by VK_EXT_pipeline_library_group_handles
typedef struct VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           pipelineLibraryGroupHandles;
} VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- <span id="features-pipelineLibraryGroupHandles"></span>
  `pipelineLibraryGroupHandles` indicates whether the implementation
  supports querying group handles directly from a ray tracing pipeline
  library, and guarantees bitwise identical group handles for such
  libraries when linked into other pipelines.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_LIBRARY_GROUP_HANDLES_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_library_group_handles](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_library_group_handles.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
