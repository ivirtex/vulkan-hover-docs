# VkPipelineExecutablePropertiesKHR(3) Manual Page

## Name

VkPipelineExecutablePropertiesKHR - Structure describing a pipeline
executable



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineExecutablePropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_pipeline_executable_properties
typedef struct VkPipelineExecutablePropertiesKHR {
    VkStructureType       sType;
    void*                 pNext;
    VkShaderStageFlags    stages;
    char                  name[VK_MAX_DESCRIPTION_SIZE];
    char                  description[VK_MAX_DESCRIPTION_SIZE];
    uint32_t              subgroupSize;
} VkPipelineExecutablePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stages` is a bitmask of zero or more
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) indicating which
  shader stages (if any) were principally used as inputs to compile this
  pipeline executable.

- `name` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a
  null-terminated UTF-8 string which is a short human readable name for
  this pipeline executable.

- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char`
  containing a null-terminated UTF-8 string which is a human readable
  description for this pipeline executable.

- `subgroupSize` is the subgroup size with which this pipeline
  executable is dispatched.

## <a href="#_description" class="anchor"></a>Description

Not all implementations have a 1:1 mapping between shader stages and
pipeline executables and some implementations **may** reduce a given
shader stage to fixed function hardware programming such that no
pipeline executable is available. No guarantees are provided about the
mapping between shader stages and pipeline executables and `stages`
**should** be considered a best effort hint. Because the application
**cannot** rely on the `stages` field to provide an exact description,
`name` and `description` provide a human readable name and description
which more accurately describes the given pipeline executable.

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineExecutablePropertiesKHR-sType-sType"
  id="VUID-VkPipelineExecutablePropertiesKHR-sType-sType"></a>
  VUID-VkPipelineExecutablePropertiesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_PROPERTIES_KHR`

- <a href="#VUID-VkPipelineExecutablePropertiesKHR-pNext-pNext"
  id="VUID-VkPipelineExecutablePropertiesKHR-pNext-pNext"></a>
  VUID-VkPipelineExecutablePropertiesKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_pipeline_executable_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_executable_properties.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutablePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineExecutablePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
