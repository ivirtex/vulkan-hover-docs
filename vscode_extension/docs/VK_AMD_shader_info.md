# VK\_AMD\_shader\_info(3) Manual Page

## Name

VK\_AMD\_shader\_info - device extension



## [](#_registered_extension_number)Registered Extension Number

43

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_special_use)Special Use

- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Jaakko Konttinen [\[GitHub\]jaakkoamd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_info%5D%20%40jaakkoamd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_info%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-10-09

**IP Status**

No known IP claims.

**Contributors**

- Jaakko Konttinen, AMD

## [](#_description)Description

This extension adds a way to query certain information about a compiled shader which is part of a pipeline. This information may include shader disassembly, shader binary and various statistics about a shader’s resource usage.

While this extension provides a mechanism for extracting this information, the details regarding the contents or format of this information are not specified by this extension and may be provided by the vendor externally.

Furthermore, all information types are optionally supported, and users should not assume every implementation supports querying every type of information.

## [](#_new_commands)New Commands

- [vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderInfoAMD.html)

## [](#_new_structures)New Structures

- [VkShaderResourceUsageAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderResourceUsageAMD.html)
- [VkShaderStatisticsInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStatisticsInfoAMD.html)

## [](#_new_enums)New Enums

- [VkShaderInfoTypeAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderInfoTypeAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_SHADER_INFO_EXTENSION_NAME`
- `VK_AMD_SHADER_INFO_SPEC_VERSION`

## [](#_examples)Examples

This example extracts the register usage of a fragment shader within a particular graphics pipeline:

```c++
extern VkDevice device;
extern VkPipeline gfxPipeline;

PFN_vkGetShaderInfoAMD pfnGetShaderInfoAMD = (PFN_vkGetShaderInfoAMD)vkGetDeviceProcAddr(
    device, "vkGetShaderInfoAMD");

VkShaderStatisticsInfoAMD statistics = {};

size_t dataSize = sizeof(statistics);

if (pfnGetShaderInfoAMD(device,
    gfxPipeline,
    VK_SHADER_STAGE_FRAGMENT_BIT,
    VK_SHADER_INFO_TYPE_STATISTICS_AMD,
    &dataSize,
    &statistics) == VK_SUCCESS)
{
    printf("VGPR usage: %d\n", statistics.resourceUsage.numUsedVgprs);
    printf("SGPR usage: %d\n", statistics.resourceUsage.numUsedSgprs);
}
```

The following example continues the previous example by subsequently attempting to query and print shader disassembly about the fragment shader:

```c++
// Query disassembly size (if available)
if (pfnGetShaderInfoAMD(device,
    gfxPipeline,
    VK_SHADER_STAGE_FRAGMENT_BIT,
    VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD,
    &dataSize,
    nullptr) == VK_SUCCESS)
{
    printf("Fragment shader disassembly:\n");

    void* disassembly = malloc(dataSize);

    // Query disassembly and print
    if (pfnGetShaderInfoAMD(device,
        gfxPipeline,
        VK_SHADER_STAGE_FRAGMENT_BIT,
        VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD,
        &dataSize,
        disassembly) == VK_SUCCESS)
    {
        printf((char*)disassembly);
    }

    free(disassembly);
}
```

## [](#_version_history)Version History

- Revision 1, 2017-10-09 (Jaakko Konttinen)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_shader_info).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0