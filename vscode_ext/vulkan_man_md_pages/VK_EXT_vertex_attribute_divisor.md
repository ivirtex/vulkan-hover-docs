# VK_EXT_vertex_attribute_divisor(3) Manual Page

## Name

VK_EXT_vertex_attribute_divisor - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

191

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_vertex_attribute_divisor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vertex_attribute_divisor.html)
  extension

## <a href="#_contact" class="anchor"></a>Contact

- Vikram Kushwaha <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_vertex_attribute_divisor%5D%20@vkushwaha%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_vertex_attribute_divisor%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>vkushwaha</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-08-03

**IP Status**  
No known IP claims.

**Contributors**  
- Vikram Kushwaha, NVIDIA

- Faith Ekstrand, Intel

## <a href="#_description" class="anchor"></a>Description

This extension allows instance-rate vertex attributes to be repeated for
certain number of instances instead of advancing for every instance when
instanced rendering is enabled.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkVertexInputBindingDivisorDescriptionEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT.html)

- Extending
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html):

  - [VkPipelineVertexInputDivisorStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_VERTEX_ATTRIBUTE_DIVISOR_EXTENSION_NAME`

- `VK_EXT_VERTEX_ATTRIBUTE_DIVISOR_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) What is the effect of a non-zero value for `firstInstance`?

**RESOLVED**: The Vulkan API should follow the OpenGL convention and
offset attribute fetching by `firstInstance` while computing vertex
attribute offsets.

2\) Should zero be an allowed divisor?

**RESOLVED**: Yes. A zero divisor means the vertex attribute is repeated
for all instances.

## <a href="#_examples" class="anchor"></a>Examples

To create a vertex binding such that the first binding uses instanced
rendering and the same attribute is used for every 4 draw instances, an
application could use the following set of structures:

``` c
    const VkVertexInputBindingDivisorDescriptionEXT divisorDesc =
    {
        .binding = 0,
        .divisor = 4
    };

    const VkPipelineVertexInputDivisorStateCreateInfoEXT divisorInfo =
    {
        .sType = VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT,
        .pNext = NULL,
        .vertexBindingDivisorCount = 1,
        .pVertexBindingDivisors = &divisorDesc
    }

    const VkVertexInputBindingDescription binding =
    {
        .binding = 0,
        .stride = sizeof(Vertex),
        .inputRate = VK_VERTEX_INPUT_RATE_INSTANCE
    };

    const VkPipelineVertexInputStateCreateInfo viInfo =
    {
        .sType = VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_CREATE_INFO,
        .pNext = &divisorInfo,
        ...
    };
    //...
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-12-04 (Vikram Kushwaha)

  - First Version

- Revision 2, 2018-07-16 (Faith Ekstrand)

  - Adjust the interaction between `divisor` and `firstInstance` to
    match the OpenGL convention.

  - Disallow divisors of zero.

- Revision 3, 2018-08-03 (Vikram Kushwaha)

  - Allow a zero divisor.

  - Add a physical device features structure to query/enable this
    feature.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_vertex_attribute_divisor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
