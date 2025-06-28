# VK\_EXT\_vertex\_attribute\_divisor(3) Manual Page

## Name

VK\_EXT\_vertex\_attribute\_divisor - device extension



## [](#_registered_extension_number)Registered Extension Number

191

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_vertex\_attribute\_divisor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_vertex_attribute_divisor.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_vertex_attribute_divisor%5D%20%40vkushwaha%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_vertex_attribute_divisor%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-08-03

**IP Status**

No known IP claims.

**Contributors**

- Vikram Kushwaha, NVIDIA
- Faith Ekstrand, Intel

## [](#_description)Description

This extension allows instance-rate vertex attributes to be repeated for certain number of instances instead of advancing for every instance when instanced rendering is enabled.

## [](#_new_structures)New Structures

- [VkVertexInputBindingDivisorDescriptionEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDivisorDescriptionEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT.html)
- Extending [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html):
  
  - [VkPipelineVertexInputDivisorStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputDivisorStateCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_VERTEX_ATTRIBUTE_DIVISOR_EXTENSION_NAME`
- `VK_EXT_VERTEX_ATTRIBUTE_DIVISOR_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_EXT`

## [](#_issues)Issues

1\) What is the effect of a non-zero value for `firstInstance`?

**RESOLVED**: The Vulkan API should follow the OpenGL convention and offset attribute fetching by `firstInstance` while computing vertex attribute offsets.

2\) Should zero be an allowed divisor?

**RESOLVED**: Yes. A zero divisor means the vertex attribute is repeated for all instances.

## [](#_examples)Examples

To create a vertex binding such that the first binding uses instanced rendering and the same attribute is used for every 4 draw instances, an application could use the following set of structures:

```c++
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

## [](#_version_history)Version History

- Revision 1, 2017-12-04 (Vikram Kushwaha)
  
  - First Version
- Revision 2, 2018-07-16 (Faith Ekstrand)
  
  - Adjust the interaction between `divisor` and `firstInstance` to match the OpenGL convention.
  - Disallow divisors of zero.
- Revision 3, 2018-08-03 (Vikram Kushwaha)
  
  - Allow a zero divisor.
  - Add a physical device features structure to query/enable this feature.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_vertex_attribute_divisor)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0